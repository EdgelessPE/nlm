package service

import (
	"fmt"
	"nlm/context"
	"os"
	"time"
)

func CreateLog(ctx context.PipelineContext, moduleName string) (*os.File, error) {
	// 创建 log 目录
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.MkdirAll("logs", 0755)
	}

	// 创建 log 文件
	file, err := os.Create(fmt.Sprintf("logs/pipeline-%s-%s.log", ctx.Id, moduleName))
	if err != nil {
		return nil, err
	}

	return file, nil
}

// 清理 30 天前的日志
func CleanLogs() error {
	println("Start cleaning outdated log files...")
	files, err := os.ReadDir("logs")
	if err != nil {
		return err
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return err
		}
		if info.ModTime().Before(time.Now().AddDate(0, 0, -30)) {
			println("Cleaning outdated log file:", file.Name())
			os.Remove(fmt.Sprintf("logs/%s", file.Name()))
		}
	}
	return nil
}
