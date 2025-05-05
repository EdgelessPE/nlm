package service

import (
	"nlm/config"
	"nlm/db"
	"nlm/model"
	"nlm/vo"
)

func GetEptToolchain() (vo.MirrorEptToolchain, error) {
	var epts []model.Ept
	if err := db.DB.Model(&model.Ept{}).Find(&epts).Error; err != nil {
		return vo.MirrorEptToolchain{}, err
	}

	releases := make([]vo.MirrorEptToolchainRelease, 0)
	for _, ept := range epts {
		url, err := GetStorageUrl(ept.StorageKey)
		if err != nil {
			return vo.MirrorEptToolchain{}, err
		}
		releases = append(releases, vo.MirrorEptToolchainRelease{
			Name:      ept.Name,
			Version:   ept.Version,
			Url:       url,
			Size:      ept.FileSize,
			Timestamp: ept.CreatedAt.Unix(),
			Integrity: ept.Integrity,
		})
	}

	return vo.MirrorEptToolchain{
		Update: vo.MirrorEptToolchainUpdate{
			WildGaps: config.ENV.EPT_WILD_GAPS,
		},
		Releases: releases,
	}, nil
}
