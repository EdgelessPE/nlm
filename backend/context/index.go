package context

import (
	"context"
)

type PipelineContext struct {
	context.Context
	BotLog string
}

func NewPipelineContext() PipelineContext {
	return PipelineContext{
		Context: context.Background(),
		BotLog:  "",
	}
}
