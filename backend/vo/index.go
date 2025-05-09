package vo

type BaseResponse[T any] struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  T      `json:"data"`
	Total int64  `json:"total"`
}
