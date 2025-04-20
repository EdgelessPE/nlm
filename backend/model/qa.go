package model

type QaResult struct {
	Base
	IsSuccess        bool   `json:"isSuccess"`
	ResultStorageKey string `json:"resultStorageKey"`

	NepId string `gorm:"index;not null"`
	Nep   *Nep   `gorm:"foreignKey:NepId;references:ID;constraint:OnDelete:CASCADE"`
}
