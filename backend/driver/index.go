package driver

import (
	"nlm/driver/downloader"
	"nlm/driver/uploader"
)

type UploadDriver interface {
	Init(targetBucketName string, rootDir string) error
	Upload(sourceFilePath string, subDir string, uuid string) error
}

type DownloadDriver interface {
	Init(entryUrl string, mountPath string) error
	GetDownloadUrl(subDir string, uuid string) (string, error)
}

var UploadDriverRegistry = map[string]UploadDriver{
	"rclone": &uploader.RcloneUploaderDriver{},
}

var DownloadDriverRegistry = map[string]DownloadDriver{
	"nginx": &downloader.NginxDownloaderDriver{},
}
