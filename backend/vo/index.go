package vo

type BaseResponse[T any] struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  T      `json:"data"`
	Total int64  `json:"total"`
}

type BasicTableParams struct {
	// 分页
	Offset int `form:"offset"`
	Limit  int `form:"limit"`

	// 搜索
	Q string `form:"q"`

	// 排序
	Sort   int    `form:"sort"`
	SortBy string `form:"sortBy"`
}
