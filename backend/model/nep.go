package model

type Nep struct {
	Base
	Scope string `gorm:"index"`
	Name  string `gorm:"index"`

	LatestReleaseVersion string
}

type Release struct {
	Base
	Version        string `gorm:"index"`
	Flags          string
	FileName       string
	StorageKey     string
	MetaStorageKey string

	NepId string `gorm:"index;not null"`
	Nep   *Nep   `gorm:"foreignKey:NepId;references:ID;constraint:OnDelete:CASCADE"`

	PipelineId string `gorm:"index"`

	IsSuccess          bool
	QaResultStorageKey string
}
