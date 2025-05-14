package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"nlm/config"
	"nlm/context"
	"nlm/db"
	"nlm/model"
	"nlm/utils"
	"nlm/vo"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

func BotGenerateDatabase() ([]model.Nep, error) {
	var neps []model.Nep

	err := db.DB.Find(&neps).Error

	if err != nil {
		return nil, err
	}

	record := make(map[string]vo.BotDatabaseNode)
	for _, nep := range neps {
		// 如果没有构建记录可以直接跳过
		if nep.LatestReleaseVersion == "" {
			continue
		}
		node := vo.BotDatabaseNode{
			Recent: vo.BotDatabaseNodeRecent{
				Health:        3,
				LatestVersion: nep.LatestReleaseVersion,
				ErrorMessage:  "",
				Builds:        []vo.BotBuildStatus{},
			},
		}
		record[nep.Scope+"_"+nep.Name] = node
		// 找出所有最新版的 release，用 flags 再赋值一次
		releases, _, err := GetReleases(vo.ReleaseParams{
			NepID:        nep.ID.String(),
			Version:      nep.LatestReleaseVersion,
			IsBotSuccess: &[]bool{true}[0],
			IsQaSuccess:  &[]bool{true}[0],
		})
		if err != nil {
			return nil, err
		}
		for _, release := range releases {
			record[nep.Scope+"_"+nep.Name+"_"+release.Flags] = node
		}
	}

	// 写到 bot 数据库文件
	text, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	os.WriteFile(config.ENV.BOT_DATABASE_FILE, text, 0644)
	log.Println("Bot database generated:", string(text))
	return neps, nil
}

func storeBuilds(ctx *context.PipelineContext, nep model.Nep, fileNames []string) ([]model.Release, error) {
	filesDir := filepath.Join(config.ENV.BOT_BUILDS_DIR, nep.Scope, nep.Name)

	// 检查 builds 目录中是否存在这些文件
	for _, fileName := range fileNames {
		if _, err := os.Stat(filepath.Join(filesDir, fileName)); os.IsNotExist(err) {
			return nil, errors.New("file not found: " + err.Error())
		}
		if _, err := os.Stat(filepath.Join(filesDir, fileName+".meta")); os.IsNotExist(err) {
			return nil, errors.New("meta file not found: " + err.Error())
		}
	}
	// 依次保存并生成结果
	var builds []model.Release
	for _, fileName := range fileNames {
		// 解析文件名
		parsed, err := utils.ParseNepFileName(fileName)
		if err != nil {
			return nil, err
		}
		// 获取文件大小
		fileStat, err := os.Stat(filepath.Join(filesDir, fileName))
		if err != nil {
			return nil, err
		}
		// 添加到 storage
		storageKey, err := AddStorage(filepath.Join(filesDir, fileName), false)
		if err != nil {
			return nil, err
		}
		// 读取 meta 内容
		metaHandler, err := os.Open(filepath.Join(filesDir, fileName+".meta"))
		if err != nil {
			return nil, err
		}
		var metaToml interface{}
		err = toml.NewDecoder(metaHandler).Decode(&metaToml)
		if err != nil {
			return nil, err
		}
		metaJson, err := json.Marshal(metaToml)
		if err != nil {
			return nil, err
		}

		b := model.Release{
			NepId:        nep.ID.String(),
			Version:      parsed.Version,
			Flags:        utils.SortFlags(parsed.Flags),
			FileName:     fileName,
			FileSize:     fileStat.Size(),
			StorageKey:   storageKey,
			Meta:         metaJson,
			PipelineId:   ctx.Id,
			IsBotSuccess: true,
		}
		db.DB.Create(&b)
		// 加载 Nep 外键
		db.DB.Model(&b).Association("Nep").Find(&b.Nep)

		// 判断上一个版本是否为新的大版本
		if nep.LatestReleaseVersion != "" && utils.GetMajorVersion(nep.LatestReleaseVersion) != utils.GetMajorVersion(parsed.Version) {
			log.Println("Marking ", parsed.Version, " as last major version for ", nep.Scope, "/", nep.Name)
			releases, _, err := GetReleases(vo.ReleaseParams{
				NepID:        nep.ID.String(),
				Version:      nep.LatestReleaseVersion,
				IsBotSuccess: &[]bool{true}[0],
				IsQaSuccess:  &[]bool{true}[0],
			})
			if err != nil {
				return nil, err
			}
			for _, release := range releases {
				log.Println("Marking release ", release.FileName)
				release.IsLastMajor = true
				db.DB.Save(&release)
			}
		}

		builds = append(builds, b)
	}
	return builds, nil
}

func BotRun(ctx *context.PipelineContext, tasks []string, force bool) ([]model.Release, error) {
	// 创建日志
	logFile, err := CreateLog(ctx, "bot")
	if err != nil {
		return nil, err
	}
	defer logFile.Close()

	// 删除 builds 目录
	os.RemoveAll(config.ENV.BOT_BUILDS_DIR)
	// 删除结果文件
	os.Remove(config.ENV.BOT_RESULT_FILE)

	// 运行 bot
	cmdSplit := strings.Split(config.ENV.BOT_RUN_CMD, " ")
	if len(tasks) > 0 {
		cmdSplit = append(cmdSplit, "-t", strings.Join(tasks, ","))
	}
	if force {
		cmdSplit = append(cmdSplit, "-f")
	}
	cmd := exec.CommandContext(ctx.Context, cmdSplit[0], cmdSplit[1:]...)
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	cmd.Dir = config.ENV.BOT_DIR
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// 读取 result.json
	result, err := os.ReadFile(config.ENV.BOT_RESULT_FILE)
	if err != nil {
		return nil, err
	}
	log.Println(string(result))
	var botResult vo.BotResult
	err = json.Unmarshal(result, &botResult)
	if err != nil {
		return nil, err
	}

	botBuilds := make([]model.Release, 0)
	for _, node := range botResult.Success {
		// 确认文件名都可以被解析
		for _, fileName := range node.FileNames {
			_, err := utils.ParseNepFileName(fileName)
			if err != nil {
				return nil, fmt.Errorf("failed to parse build's file name '%s' : %s", fileName, err.Error())
			}
		}

		// 获取 NepId
		nep, err := GetNep(node.Scope, utils.CleanBotTaskName(node.TaskName))
		if err != nil {
			return nil, err
		}

		// 保存 builds
		b, err := storeBuilds(ctx, nep, node.FileNames)
		if err != nil {
			return nil, err
		}
		botBuilds = append(botBuilds, b...)
	}
	for _, node := range botResult.Failed {
		nep, err := GetNep(node.Scope, utils.CleanBotTaskName(node.TaskName))
		if err != nil {
			return nil, err
		}
		// 写入失败的 release 记录
		db.DB.Create(&model.Release{
			NepId:        nep.ID.String(),
			PipelineId:   ctx.Id,
			IsBotSuccess: false,
			BotErrMsg:    node.ErrorMessage,
		})
	}

	return botBuilds, nil
}
