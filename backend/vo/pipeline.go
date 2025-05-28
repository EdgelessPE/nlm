package vo

type PipelineParams struct {
	BasicTableParams
	ModelName string `form:"model_name"`
}
