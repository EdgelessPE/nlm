package service

import (
	"fmt"
	"nlm/context"
	"os"
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
