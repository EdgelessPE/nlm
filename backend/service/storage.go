package service

import (
	"fmt"
	"nlm/config"
	"nlm/db"
	"nlm/model"
	"nlm/vo"
	"os"
	"path/filepath"
	"time"
)

func AddStorage(sourceFilePath string, syncLocation []vo.SyncLocation) (string, error) {
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

	// TODO:调度文件同步任务
	go func() {
		for _, location := range syncLocation {
			switch location {
			case vo.HomeServer:
				break
			case vo.Cloud189:
				break
			case vo.Quark:
				break
			case vo.KanuoCloud:
				break
			}
		}

		// 更新同步状态
		db.DB.Model(&s).Update("sync_finished_at", time.Now())
	}()

	return uuid, nil
}
