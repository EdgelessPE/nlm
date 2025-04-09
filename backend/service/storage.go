package service

import (
	"fmt"
	"nlm/config"
	"nlm/db"
	"nlm/model"
	storage_drivers "nlm/service/storage-drivers"
	"os"
	"path/filepath"
	"time"

	"github.com/cespare/cp"
)

func syncFile(uuid string, syncToExpensiveStorage bool) error {
	sourceFilePath := filepath.Join(config.ENV.STORAGE_TEMP_DIR, uuid)
	if _, err := os.Stat(sourceFilePath); os.IsNotExist(err) {
		return fmt.Errorf("sync error: source file not found: %s", sourceFilePath)
	}

	storageConfig := config.ENV.STORAGE_CONFIG
	for _, config := range storageConfig {
		if !syncToExpensiveStorage && config.Expensive {
			continue
		}

		driver := storage_drivers.Registry[config.Driver]
		err := driver.Init(config.StorageName, config.BaseDir)
		if err != nil {
			return err
		}

		err = driver.Upload(uuid, sourceFilePath)
		if err != nil {
			return err
		}
	}

	return nil
}

func AddStorage(sourceFilePath string, syncToExpensiveStorage bool) (string, error) {
	// 存库并分配 UUID
	var s model.Storage
	s.FileName = filepath.Base(sourceFilePath)
	db.DB.Create(&s)

	// 获取 UUID
	uuid := fmt.Sprint(s.ID)

	// 移动到临时目录
	tempDir := config.ENV.STORAGE_TEMP_DIR
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", err
	}
	tempFilePath := filepath.Join(tempDir, uuid)
	err := os.Rename(sourceFilePath, tempFilePath)
	if err != nil {
		return "", err
	}

	// 调度文件同步任务
	go func() {
		fmt.Println("start syncing file", uuid)
		err := syncFile(uuid, syncToExpensiveStorage)
		if err != nil {
			fmt.Println("sync error: ", err)
		}
		// 更新同步状态
		db.DB.Model(&s).Update("sync_finished_at", time.Now())
		fmt.Println("sync finished", uuid)
	}()

	return uuid, nil
}

func FetchStorage(uuid string, toDir string) (string, error) {
	// 查询文件名
	var s model.Storage
	db.DB.Where("id = ?", uuid).First(&s)
	if s.FileName == "" {
		return "", fmt.Errorf("can't found storage for uuid: %s", uuid)
	}

	// 从临时存储中获取文件
	tempDir := config.ENV.STORAGE_TEMP_DIR
	tempFilePath := filepath.Join(tempDir, uuid)
	if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
		return "", fmt.Errorf("can't found temp file for uuid: %s", uuid)
	}

	// 复制文件到目标位置
	targetFilePath := filepath.Join(toDir, s.FileName)
	if err := os.MkdirAll(toDir, 0755); err != nil {
		return "", err
	}
	if err := cp.CopyFile(tempFilePath, targetFilePath); err != nil {
		return "", err
	}

	return targetFilePath, nil
}

func GetStorageUrl(uuid string) (string, error) {
	storageConfig := config.ENV.STORAGE_CONFIG
	for _, config := range storageConfig {
		if config.PublicUrlBase == "" {
			continue
		}
		driver := storage_drivers.Registry[config.Driver]
		exists, err := driver.Exists(uuid)
		if err != nil {
			return "", err
		}
		if exists {
			return filepath.Join(config.PublicUrlBase, uuid), nil
		}
	}
	return "", fmt.Errorf("can't found storage for uuid: %s", uuid)
}
