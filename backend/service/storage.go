package service

import (
	"fmt"
	"log"
	"nlm/config"
	"nlm/db"
	"nlm/driver"
	"nlm/model"
	"nlm/utils"
	"os"
	"path/filepath"
	"time"

	"github.com/alitto/pond/v2"
	"github.com/cespare/cp"
)

func syncFile(uuid string) error {
	sourceFilePath := filepath.Join(config.ENV.STORAGE_TEMP_DIR, uuid)
	if _, err := os.Stat(sourceFilePath); os.IsNotExist(err) {
		return fmt.Errorf("sync error: source file not found: %s", sourceFilePath)
	}

	storageConfig := config.ENV.STORAGE_CONFIG
	for _, config := range storageConfig {
		driver := driver.UploadDriverRegistry[config.UploaderDriver]
		err := driver.Init(config.UploaderTargetBucketName, config.UploaderRootDir)
		if err != nil {
			return err
		}

		err = driver.Upload(sourceFilePath, utils.GetUUIDSubDir(uuid), uuid)
		if err != nil {
			return err
		}
	}

	return nil
}

var pool = pond.NewPool(3)

func AddStorage(sourceFilePath string, compressWithZstd bool) (string, error) {
	// 获取文件大小
	fileStat, err := os.Stat(sourceFilePath)
	if err != nil {
		return "", err
	}

	// 存库并分配 UUID
	var s model.Storage
	s.FileName = filepath.Base(sourceFilePath)
	s.FileSize = fileStat.Size()
	s.Compressed = compressWithZstd
	db.DB.Create(&s)

	// 获取 UUID
	uuid := fmt.Sprint(s.ID)

	// 移动到临时目录
	tempDir := config.ENV.STORAGE_TEMP_DIR
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		return "", err
	}
	tempFilePath := filepath.Join(tempDir, uuid)
	if compressWithZstd {
		// 使用 zstd 压缩
		err := utils.CompressFileWithZstd(sourceFilePath, tempFilePath)
		if err != nil {
			return "", err
		}
	} else {
		err := os.Rename(sourceFilePath, tempFilePath)
		if err != nil {
			return "", err
		}
	}

	// 调度文件同步任务
	pool.Submit(func() {
		log.Println("start syncing file", uuid)
		err := syncFile(uuid)
		if err != nil {
			log.Println("sync error: ", err)
		}
		// 更新同步状态
		db.DB.Model(&s).Update("sync_finished_at", time.Now())
		log.Println("sync finished", uuid)
	})

	return uuid, nil
}

func FetchStorage(uuid string, toDir string) (string, error) {
	fmt.Printf("fetching storage %s to %s\n", uuid, toDir)
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
	if s.Compressed {
		err := utils.DecompressFileWithZstd(tempFilePath, targetFilePath)
		if err != nil {
			return "", err
		}
	} else {
		if err := cp.CopyFile(targetFilePath, tempFilePath); err != nil {
			return "", err
		}
	}

	return targetFilePath, nil
}

func GetStorageUrl(uuid string) (string, error) {
	storageConfig := config.ENV.STORAGE_CONFIG
	for _, config := range storageConfig {
		if config.DownloaderEntryUrl == "" {
			continue
		}
		driver := driver.DownloadDriverRegistry[config.DownloaderDriver]
		err := driver.Init(config.DownloaderEntryUrl, config.DownloaderMountPath)
		if err != nil {
			return "", err
		}
		downloadUrl, err := driver.GetDownloadUrl(utils.GetUUIDSubDir(uuid), uuid)
		if err != nil {
			return "", err
		}
		return downloadUrl, nil
	}
	return "", fmt.Errorf("can't found storage for uuid: %s", uuid)
}

func DeleteStorage(uuid string) {
	storageConfig := config.ENV.STORAGE_CONFIG
	for _, config := range storageConfig {
		driver := driver.UploadDriverRegistry[config.UploaderDriver]
		err := driver.Delete(utils.GetUUIDSubDir(uuid), uuid)
		if err != nil {
			log.Println("Warning: failed to delete storage: ", err.Error())
		}
	}
}

// 清理 30天前的临时存储文件
func CleanTempStorage() error {
	return utils.CleanOutdatedFiles(config.ENV.STORAGE_TEMP_DIR)
}
