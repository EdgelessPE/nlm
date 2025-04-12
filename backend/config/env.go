package config

import (
	"log"
	"nlm/vo"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Env struct {
	DB_HOST     string
	DB_PORT     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	BOT_DIR           string
	BOT_TASKS_DIR     string
	BOT_DATABASE_FILE string
	BOT_RUN_CMD       string
	BOT_RESULT_FILE   string
	BOT_BUILDS_DIR    string
	STORAGE_TEMP_DIR  string
	STORAGE_CONFIG    []vo.StorageConfig
}

var ENV Env

func init() {
	// 读取 .env.toml 文件
	envFile, err := os.Open(".env.toml")
	if err != nil {
		log.Fatal(err)
	}
	defer envFile.Close()

	// 解析 .env.toml 文件
	err = toml.NewDecoder(envFile).Decode(&ENV)
	if err != nil {
		log.Fatal(err)
	}
}
