package model

type Storage struct {
	Base
	SourceFilePath string `gorm:"not null"`
}
