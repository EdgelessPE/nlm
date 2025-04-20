package service

import (
	"fmt"
	"nlm/config"
	"nlm/context"
	"nlm/db"
	"nlm/model"
	"os"
	"os/exec"
	"path/filepath"
)

func QaPreparePackages(builds []model.Release) error {
	// 清理存储和报告目录
	if err := os.RemoveAll(config.ENV.QA_STORAGE_DIR); err != nil {
		return err
	}
	if err := os.RemoveAll(config.ENV.QA_REPORTS_DIR); err != nil {
		return err
	}

	for _, build := range builds {
		// 创建目录
		dir := filepath.Join(config.ENV.QA_STORAGE_DIR, build.Nep.Scope, build.Nep.Name)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		// 获取构建
		_, err := FetchStorage(build.StorageKey, dir)
		if err != nil {
			return err
		}
	}

	return nil
}

func QaRun(ctx context.PipelineContext, builds []model.Release) ([]model.Release, error) {
	// 创建日志
	logFile, err := CreateLog(ctx, "qa")
	if err != nil {
		return nil, err
	}
	defer logFile.Close()

	// 删除 reports 目录
	if err := os.RemoveAll(config.ENV.QA_REPORTS_DIR); err != nil {
		return nil, err
	}

	// 运行 qa
	cmd := exec.CommandContext(ctx.Context, config.ENV.QA_RUN_CMD, config.ENV.QA_STORAGE_DIR)
	cmd.Stdout = logFile
	cmd.Stderr = logFile
	cmd.Dir = config.ENV.QA_DIR
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// 读取报告
	for _, build := range builds {
		fileName := build.FileName
		reportDir := filepath.Join(config.ENV.QA_REPORTS_DIR, build.Nep.Scope, build.Nep.Name, fileName)
		// 更新闭包
		updateBuild := func(filePath string, isSuccess bool) error {
			key, err := AddStorage(filePath, false, true)
			if err != nil {
				return err
			}
			build.QaResultStorageKey = key
			build.IsSuccess = isSuccess
			db.DB.Model(&build).Updates(map[string]interface{}{
				"qa_result_storage_key": key,
				"is_success":            isSuccess,
			})
			return nil
		}
		// 检查目录下的文件
		failedFile := filepath.Join(reportDir, "Error.txt")
		if stat, _ := os.Stat(failedFile); stat != nil {
			if err := updateBuild(failedFile, false); err != nil {
				return nil, err
			}
			continue
		}
		readmeFile := filepath.Join(reportDir, "README.md")
		if stat, _ := os.Stat(readmeFile); stat != nil {
			if err := updateBuild(readmeFile, true); err != nil {
				return nil, err
			}
			continue
		}
		return nil, fmt.Errorf("can't found report for build: %s", build.StorageKey)
	}

	return builds, nil
}
