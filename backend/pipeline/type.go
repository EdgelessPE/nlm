package pipeline

import "nlm/context"

type PipelineCreateResult struct {
	PipelineContext context.PipelineContext
	IsNewPipeline   bool
}
