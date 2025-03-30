package service

import (
	"encoding/json"
	"log"
	"nlm/context"
	"nlm/db"
	"nlm/model"
	"nlm/vo"
	"os"
	"os/exec"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

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
	os.WriteFile(os.Getenv("BOT_DATABASE_FILE"), text, 0644)
	return neps, nil
}

func BotRun(ctx context.PipelineContext) (vo.BotResult, error) {
	// 运行 bot
	cmdSplit := strings.Split(os.Getenv("BOT_RUN_CMD"), " ")
	cmd := exec.Command(cmdSplit[0], cmdSplit[1:]...)
	cmd.Dir = os.Getenv("BOT_DIR")
	output, err := cmd.Output()
	if err != nil {
		return vo.BotResult{}, err
	}

	// 将日志写入 context
	ctx.BotLog = string(output)

	// 读取 result.json
	result, err := os.ReadFile(os.Getenv("BOT_RESULT_FILE"))
	if err != nil {
		return vo.BotResult{}, err
	}
	var botResult vo.BotResult
	err = json.Unmarshal(result, &botResult)
	if err != nil {
		return vo.BotResult{}, err
	}

	return botResult, nil
}
