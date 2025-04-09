package storage_drivers

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

type RcloneDriver struct {
	targetStorageName string
	targetDir         string
}

func (d *RcloneDriver) Init(targetStorageName string, targetDir string) error {
	// 检查 rclone 是否安装
	if _, err := exec.LookPath("rclone"); err != nil {
		log.Fatal("rclone 未安装")
	}

	d.targetStorageName = targetStorageName
	d.targetDir = targetDir

	return nil
}

func (d *RcloneDriver) Upload(uuid string, sourceFilePath string) error {
	// 执行 rclone 命令
	fmt.Println("Running rclone copy", sourceFilePath, d.targetStorageName+":"+filepath.Join(d.targetDir, uuid))
	cmd := exec.Command("rclone", "copy", sourceFilePath, d.targetStorageName+":"+filepath.Join(d.targetDir, uuid))
	return cmd.Run()
}

func (d *RcloneDriver) Exists(uuid string) (bool, error) {
	cmd := exec.Command("rclone", "ls", d.targetStorageName+":"+d.targetDir)
	output, err := cmd.Output()
	if err != nil {
		return false, err
	}

	return strings.Contains(string(output), uuid), nil
}
