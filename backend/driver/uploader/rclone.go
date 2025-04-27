package uploader

import (
	"log"
	"os/exec"
	"path/filepath"
)

type RcloneUploaderDriver struct {
	targetStorageName string
	targetDir         string
}

func (d *RcloneUploaderDriver) Init(targetBucketName string, rootDir string) error {
	// 检查 rclone 是否安装
	if _, err := exec.LookPath("rclone"); err != nil {
		log.Fatal("rclone 未安装")
	}

	d.targetStorageName = targetBucketName
	d.targetDir = rootDir

	return nil
}

func (d *RcloneUploaderDriver) Upload(sourceFilePath string, subDir string, uuid string) error {
	// 执行 rclone 命令
	log.Println("Running rclone copyto", sourceFilePath, d.targetStorageName+":"+filepath.Join(d.targetDir, subDir, uuid))
	cmd := exec.Command("rclone", "copyto", sourceFilePath, d.targetStorageName+":"+filepath.Join(d.targetDir, subDir, uuid))
	return cmd.Run()
}

// func (d *RcloneUploadDriver) Exists(uuid string) (bool, error) {
// 	cmd := exec.Command("rclone", "ls", d.targetStorageName+":"+d.targetDir)
// 	output, err := cmd.Output()
// 	if err != nil {
// 		return false, err
// 	}

// 	return strings.Contains(string(output), uuid), nil
// }
