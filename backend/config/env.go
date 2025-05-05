package config

import (
	"log"
	"nlm/vo"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Env struct {
	ROOT_URL string

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

	QA_DIR         string
	QA_STORAGE_DIR string
	QA_REPORTS_DIR string
	QA_RUN_CMD     string

	MIRROR_HELLO_NAME        string
	MIRROR_HELLO_LOCALE      string
	MIRROR_HELLO_DESCRIPTION string
	MIRROR_HELLO_MAINTAINER  string
	MIRROR_HELLO_PROPERTY    vo.MirrorHelloProperty

	WEBHOOK_TOKEN string

	GITHUB_TOKEN string
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
