package context

import (
	"context"

	"github.com/google/uuid"
)

type PipelineContext struct {
	context.Context
	Id string
}

func NewPipelineContext() PipelineContext {
	return PipelineContext{
		Context: context.Background(),
		Id:      uuid.New().String(),
	}
}
