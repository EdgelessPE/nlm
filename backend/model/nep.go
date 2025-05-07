package model

import "nlm/db"

type Nep struct {
	Base
	Scope string `gorm:"index"`
	Name  string `gorm:"index"`

	LatestReleaseVersion string
}

type Release struct {
	Base
	Version    string `gorm:"index"`
	Flags      string
	FileName   string
	FileSize   int64
	StorageKey string
	Meta       db.JSON `gorm:"type:jsonb"`

	// 是否是最后一个大版本
	IsLastMajor bool `gorm:"default:false"`

	NepId string `gorm:"index;not null"`
	Nep   *Nep   `gorm:"foreignKey:NepId;references:ID;constraint:OnDelete:CASCADE"`

	PipelineId string `gorm:"index"`

	IsBotSuccess bool `gorm:"default:false"`
	IsQaSuccess  bool `gorm:"default:false"`

	QaResultStorageKey string
}
