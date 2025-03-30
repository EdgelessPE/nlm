package service

import (
	"encoding/json"
	"log"
	"nlm/db"
	"nlm/model"
	"nlm/vo"
	"os"

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
