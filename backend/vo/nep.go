package vo

type NepParams struct {
	BasicTableParams
	Scope          string `form:"scope"`
	UpdatedAtStart int64  `form:"updated_at_start"`
	UpdatedAtEnd   int64  `form:"updated_at_end"`
}

type ReleaseParams struct {
	BasicTableParams
	NepID          string `form:"nep_id" binding:"required"`
	IsBotSuccess   *bool  `form:"is_bot_success"`
	IsQaSuccess    *bool  `form:"is_qa_success"`
	Version        string `form:"version"`
	Flags          string `form:"flags"`
	CreatedAtStart int64  `form:"created_at_start"`
	CreatedAtEnd   int64  `form:"created_at_end"`
}
