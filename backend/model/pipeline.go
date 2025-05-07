package model

import "time"

type Pipeline struct {
	Base
	ModelName  string
	FinishedAt time.Time
	// running, success, failed
	Status string
	ErrMsg string
	Stage  string
}
