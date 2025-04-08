package utils

import (
	"fmt"
	"log"
	"os/exec"
)

func init() {
	// 检查 rclone 是否安装
	if _, err := exec.LookPath("rclone"); err != nil {
		log.Fatal("rclone 未安装")
	}
}

func RcloneCp(sourceFilePath string, targetStorageName string, targetDir string) error {
	// 执行 rclone 命令
	fmt.Println("Running rclone copy", sourceFilePath, targetStorageName+":"+targetDir)
	cmd := exec.Command("rclone", "copy", sourceFilePath, targetStorageName+":"+targetDir)
	return cmd.Run()
}
