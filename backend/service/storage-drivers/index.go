package storage_drivers

// 存储驱动枚举（rclone、本地）
type StorageDriverEnum int

const (
	StorageDriverEnumRclone StorageDriverEnum = iota + 1
	// 软件官方客户端，使用同步功能
	StorageDriverEnumOfficialClient
)

type StorageDriver interface {
	Init(targetStorageName string, targetDir string) error
	Upload(uuid string, sourceFilePath string) error
	Exists(uuid string) (bool, error)
}

var Registry = map[StorageDriverEnum]StorageDriver{
	StorageDriverEnumRclone:         &RcloneDriver{},
	StorageDriverEnumOfficialClient: &VoidDriver{},
}
