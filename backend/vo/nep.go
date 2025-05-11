package vo

type NepParams struct {
	// 分页
	Offset int `form:"offset"`
	Limit  int `form:"limit"`

	// 搜索
	Q string `form:"q"`

	// 排序
	Sort  string `form:"sort"`
	Order string `form:"order"`

	Scope string `form:"scope"`
}
