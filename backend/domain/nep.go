package domain

import (
	"log"
	"nlm/service"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func InitNepsWithBotTask() {
	tasks_dir := os.Getenv("BOT_TASKS_DIR")
	scope_dirs, err := os.ReadDir(tasks_dir)
	if err != nil {
		log.Fatal("Error reading bot tasks directory: " + err.Error())
	}

	for _, scope_dir := range scope_dirs {
		name_dirs, err := os.ReadDir(filepath.Join(tasks_dir, scope_dir.Name()))
		if err != nil {
			log.Fatal("Error reading bot tasks directory: " + err.Error())
		}
		for _, name_dir := range name_dirs {
			if name_dir.IsDir() {
				// 读取 config.toml 文件
				text, err := os.ReadFile(filepath.Join(tasks_dir, scope_dir.Name(), name_dir.Name(), "config.toml"))
				if err != nil {
					log.Fatal("Error reading config file: " + err.Error())
				}

				// 解析 config.toml 文件
				var config map[string]interface{}
				err = toml.Unmarshal(text, &config)
				if err != nil {
					log.Fatal("Error parsing config file: " + err.Error())
				}

				// 获取 name 和 scope
				name := config["task"].(map[string]interface{})["name"].(string)
				scope := config["task"].(map[string]interface{})["scope"].(string)

				// 创建数据库
				d, err := service.AddNep(scope, name)
				if err == nil {
					println("Created nep: "+d.ID.String(), "scope: "+scope, "name: "+name)
				}
			}
		}
	}

}
