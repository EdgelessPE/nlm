package model

import (
	"time"
)

type Nep struct {
	Base
	Scope    string
	Name     string
	Releases []Release
}

type Release struct {
	Base
	Version string
	Flags   []string

	PutawayAt  time.Time
	PipelineId string
}
