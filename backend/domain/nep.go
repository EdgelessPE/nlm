package domain

import (
	"log"
	"nlm/config"
	"nlm/service"
	"nlm/utils"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

func InitNepsWithBotTask() {
	tasks_dir := config.ENV.BOT_TASKS_DIR
	scope_dirs, err := os.ReadDir(tasks_dir)
	if err != nil {
		log.Fatal("Error reading bot tasks directory: " + err.Error())
	}

	for _, scope_dir := range scope_dirs {
		if !scope_dir.IsDir() {
			continue
		}
		name_dirs, err := os.ReadDir(filepath.Join(tasks_dir, scope_dir.Name()))
		if err != nil {
			log.Fatal("Error reading bot tasks directory: " + err.Error())
		}
		for _, name_dir := range name_dirs {
			if name_dir.IsDir() {
				config_path := filepath.Join(tasks_dir, scope_dir.Name(), name_dir.Name(), "config.toml")
				if _, err := os.Stat(config_path); os.IsNotExist(err) {
					continue
				}
				// 读取 config.toml 文件
				text, err := os.ReadFile(config_path)
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
				name := utils.CleanBotTaskName(config["task"].(map[string]interface{})["name"].(string))
				scope := config["task"].(map[string]interface{})["scope"].(string)

				// 创建数据库
				d, err := service.AddNep(scope, name)
				if err == nil {
					log.Println("Created nep: "+d.ID.String(), "scope: "+scope, "name: "+name)
				}
			}
		}
	}

}
