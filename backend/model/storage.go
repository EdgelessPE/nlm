package model

import "time"

type Storage struct {
	Base
	FileName       string `gorm:"not null"`
	FileSize       int64
	SyncFinishedAt time.Time
	Compressed     bool `gorm:"default:false"`
}
