package vo

import storage_drivers "nlm/service/storage-drivers"

// 存储配置
type StorageConfig struct {
	// 存储唯一标识
	Key string
	// 昂贵存储
	Expensive bool
	// 存储驱动
	Driver storage_drivers.StorageDriverEnum
	// 存储名称
	StorageName string
	// 存储根目录
	BaseDir string
	// 公开可访问 URL 拼接前缀
	PublicUrlBase string
}
