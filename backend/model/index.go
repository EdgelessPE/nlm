package model

import (
	"time"
)

type Nep struct {
	Base
	Scope string
	Name  string
}

type Release struct {
	Base
	NepId   string
	Version string
	Flags   string

	PutawayAt  time.Time
	PipelineId string
}
