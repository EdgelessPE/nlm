package model

import "time"

type Storage struct {
	Base
	FileName       string `gorm:"not null"`
	SyncFinishedAt time.Time
}
