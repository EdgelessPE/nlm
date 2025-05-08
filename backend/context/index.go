package context

import (
	"context"

	"github.com/google/uuid"
)

type PipelineContext struct {
	context.Context
	Id     string
	Cancel context.CancelFunc
}

func NewPipelineContext() PipelineContext {
	ctx, cancel := context.WithCancel(context.Background())
	return PipelineContext{
		Context: ctx,
		Id:      uuid.New().String(),
		Cancel:  cancel,
	}
}
