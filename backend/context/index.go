package context

import (
	"context"
)

type PipelineContext struct {
	context.Context
	BotLog string
}
