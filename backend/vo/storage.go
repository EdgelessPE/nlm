package vo

// 存储驱动枚举（rclone、本地）
type StorageDriver int

const (
	StorageDriverRclone StorageDriver = iota + 1
	// 软件官方客户端，使用同步功能
	StorageDriverOfficialClient
)

// 存储配置
type StorageConfig struct {
	// 存储唯一标识
	Key string
	// 昂贵存储
	Expensive bool
	// 存储驱动
	Driver StorageDriver
	// 存储名称
	StorageName string
	// 存储根目录
	BaseDir string
	// 公开可访问 URL 拼接前缀
	PublicUrlBase string
}
