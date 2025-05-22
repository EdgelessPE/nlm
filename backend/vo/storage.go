package vo

// 存储配置
type StorageConfig struct {
	// 存储唯一标识
	Key string

	// 上传驱动
	UploaderDriver string
	// 上传存储名称
	UploaderTargetBucketName string
	// 上传根目录
	UploaderRootDir string

	// 下载驱动
	DownloaderDriver string
	// 下载入口 URL
	DownloaderEntryUrl string
	// 下载挂载路径
	DownloaderMountPath string
}

type GetStoragesParams struct {
	BasicTableParams
	IsCompressed bool
}
