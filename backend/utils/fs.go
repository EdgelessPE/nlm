package utils

import (
	"os"
	"path/filepath"
	"time"
)

func CleanOutdatedFiles(dir string) error {
	println("Start cleaning outdated files in", dir)
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		info, err := file.Info()
		if err != nil {
			continue
		}
		if info.ModTime().Before(time.Now().AddDate(0, 0, -30)) {
			filePath := filepath.Join(dir, file.Name())
			println("Cleaning outdated file:", filePath)
			os.Remove(filePath)
		}
	}
	return nil
}
