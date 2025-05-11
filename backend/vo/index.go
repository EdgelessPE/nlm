package vo

type BaseResponse[T any] struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  T      `json:"data"`
	Total int64  `json:"total"`
}

type BasicTableParams struct {
	// 分页
	Offset int `json:"offset"`
	Limit  int `json:"limit"`

	// 搜索
	Q string `json:"q"`

	// 排序
	Sort  string `json:"sort"`
	Order string `json:"order"`
}
