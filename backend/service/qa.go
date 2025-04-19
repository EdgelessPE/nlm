package service

import (
	"nlm/config"
	"nlm/context"
	"nlm/vo"
	"os"
	"os/exec"
	"path/filepath"
)

func QaPreparePackages(builds []vo.BotBuild) error {
	// 清理存储目录
	if err := os.RemoveAll(config.ENV.QA_STORAGE_DIR); err != nil {
		return err
	}

	for _, build := range builds {
		// 创建目录
		dir := filepath.Join(config.ENV.QA_STORAGE_DIR, build.Scope, build.TaskName)
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

func QaRun(ctx context.PipelineContext) ([]vo.QaResult,error) {
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
		return nil, 	err
	}

	// 读取 reports 目录
	scopeEntries, err := os.ReadDir(config.ENV.QA_REPORTS_DIR)
	if err != nil {
		return nil, err
	}
	qaReports := make([]vo.QaResult, 0)
	for _, scopeEntry := range scopeEntries {
		scopeName := scopeEntry.Name()
		scopeDir := filepath.Join(config.ENV.QA_REPORTS_DIR, scopeName)
		taskEntries, err := os.ReadDir(scopeDir)
		if err != nil {
			return nil, err
		}
		for _, taskEntry := range taskEntries {
			taskName := taskEntry.Name()
			taskDir := filepath.Join(scopeDir, taskName)
			reportEntries, err := os.ReadDir(taskDir)
			if err != nil {
				return nil, err
			}
			for _, reportEntry := range reportEntries {
				packageName := reportEntry.Name()
				// 检查目录下的文件
				failedFile := filepath.Join(taskDir, packageName, "Error.txt")
				if stat, _ := os.Stat(failedFile); stat != nil {
					key, err := AddStorage(failedFile, false,true)
					if err != nil {
						return nil, 	err
					}
					qaReports = append(qaReports, vo.QaResult{
						Scope:            scopeName,
						TaskName:         taskName,
						IsSuccess:        false,
						ResultStorageKey: key,
					})
					continue
				}
				readmeFile := filepath.Join(taskDir, packageName, "README.md")
				if stat, _ := os.Stat(readmeFile); stat != nil {
					key, err := AddStorage(readmeFile, false,true)
					if err != nil {
						return nil, err
					}
					qaReports = append(qaReports, vo.QaResult{
						Scope:            scopeName,
						TaskName:         taskName,
						IsSuccess:        true,
						ResultStorageKey: key,
					})
					continue
				}
				println("warning: empty qa report folder", scopeName, taskName, packageName)
			}
		}
	}

	return qaReports, nil
}
