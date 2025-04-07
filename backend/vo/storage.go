package vo

// 同步位置枚举（HomeServer、天翼网盘、夸克网盘、卡诺云）
type SyncLocation int

const (
	HomeServer SyncLocation = iota + 1
	Cloud189
	Quark
	KanuoCloud
)

func (s SyncLocation) String() string {
	return []string{"HomeServer", "Cloud189", "Quark", "KanuoCloud"}[s-1]
}
