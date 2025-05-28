package vo

type PipelineParams struct {
	BasicTableParams
	ModelName string `form:"model_name"`
	Status    string `form:"status"`
}
