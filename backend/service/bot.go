package service

import (
	"encoding/json"
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
		record[nep.Scope+"_"+nep.Name] = vo.BotDatabaseNode{
			Recent: vo.BotDatabaseNodeRecent{
				Health:        3,
				LatestVersion: nep.LatestReleaseVersion,
				ErrorMessage:  "",
				Builds:        []vo.BotBuildStatus{},
			},
		}
	}

	// 写到 bot 数据库文件
	text, err := json.Marshal(record)
	if err != nil {
		return nil, err
	}
	os.WriteFile(config.ENV.BOT_DATABASE_FILE, text, 0644)
	return neps, nil
}

func storeBuilds(scope string, name string, fileNames []string) ([]vo.BotBuild, error) {
	filesDir := filepath.Join(config.ENV.BOT_BUILDS_DIR, scope, name)

	// 检查 builds 目录中是否存在这些文件
	for _, fileName := range fileNames {
		if _, err := os.Stat(filepath.Join(filesDir, fileName)); os.IsNotExist(err) {
			return nil, err
		}
		if _, err := os.Stat(filepath.Join(filesDir, fileName+".meta")); os.IsNotExist(err) {
			return nil, err
		}
	}
	// 依次保存并生成结果
	var builds []vo.BotBuild
	for _, fileName := range fileNames {
		parsed, err := utils.ParseNepFileName(fileName)
		if err != nil {
			return nil, err
		}
		storageKey, err := AddStorage(filepath.Join(filesDir, fileName), false)
		if err != nil {
			return nil, err
		}
		metaStorageKey, err := AddStorage(filepath.Join(filesDir, fileName+".meta"), true)
		if err != nil {
			return nil, err
		}
		builds = append(builds, vo.BotBuild{
			Version:        parsed.Version,
			Flags:          parsed.Flags,
			FileName:       fileName,
			StorageKey:     storageKey,
			MetaStorageKey: metaStorageKey,
		})
	}
	return builds, nil
}

func BotRun(ctx context.PipelineContext, tasks []string, force bool) (vo.BotResult, error) {
	// 创建日志
	logFile, err := CreateLog(ctx, "bot")
	if err != nil {
		return vo.BotResult{}, err
	}
	defer logFile.Close()

	// 删除 builds 目录
	os.RemoveAll(config.ENV.BOT_BUILDS_DIR)

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
		return vo.BotResult{}, err
	}

	// 读取 result.json
	result, err := os.ReadFile(config.ENV.BOT_RESULT_FILE)
	if err != nil {
		return vo.BotResult{}, err
	}
	var botResult vo.BotResult
	err = json.Unmarshal(result, &botResult)
	if err != nil {
		return vo.BotResult{}, err
	}

	for _, node := range botResult.Success {
		// 保存 builds
		b, err := storeBuilds(node.Scope, node.TaskName, node.FileNames)
		if err != nil {
			return vo.BotResult{}, err
		}

		// 获取 NepId
		nep, err := GetNep(node.Scope, node.TaskName)
		if err != nil {
			return vo.BotResult{}, err
		}

		// 保存 builds 到数据库
		for _, build := range b {
			db.DB.Create(&model.Release{
				Version:        build.Version,
				Flags:          build.Flags,
				FileName:       build.FileName,
				StorageKey:     build.StorageKey,
				MetaStorageKey: build.MetaStorageKey,
				NepId:          nep.ID.String(),
				PipelineId:     ctx.Id,
			})
		}
	}

	return botResult, nil
}
