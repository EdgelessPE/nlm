package vo

type NepParams struct {
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
	Q      string `form:"q"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`

	Scope string `form:"scope"`
}

type ReleaseParams struct {
	Offset int    `form:"offset"`
	Limit  int    `form:"limit"`
	Q      string `form:"q"`
	Sort   string `form:"sort"`
	Order  string `form:"order"`

	Scope     string `form:"scope"`
	Name      string `form:"name"`
	IsSuccess bool   `form:"is_success"`
	Version   string `form:"version"`
}
