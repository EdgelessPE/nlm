package model

type Nep struct {
	Base
	Scope           string
	Name            string
	LatestReleaseId string
	LatestRelease   *Release `gorm:"foreignKey:LatestReleaseId;references:ID"`
}

type Release struct {
	Base
	Version  string
	Flags    string
	FileName string

	NepId string
	Nep   *Nep `gorm:"foreignKey:NepId;references:ID"`

	PipelineId string
}
