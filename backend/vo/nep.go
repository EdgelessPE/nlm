package vo

type NepParams struct {
	BasicTableParams
	Scope string `form:"scope"`
}

type ReleaseParams struct {
	BasicTableParams
	NepID        string `form:"nep_id" binding:"required"`
	IsBotSuccess bool   `form:"is_bot_success"`
	IsQaSuccess  bool   `form:"is_qa_success"`
	Version      string `form:"version"`
}
