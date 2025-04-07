package model

import "gorm.io/gorm"

type Storage struct {
	gorm.Model
	SourceFilePath string `gorm:"not null"`
}
