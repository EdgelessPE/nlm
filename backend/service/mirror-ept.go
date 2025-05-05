package service

import (
	"log"
	"nlm/config"
	"nlm/db"
	"nlm/model"
	"nlm/vo"
	"time"
)

var mirrorEptToolchainCache vo.MirrorEptToolchain

func generateMirrorEptToolchain() (vo.MirrorEptToolchain, error) {
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

var lastRefreshTimeEpt time.Time = time.Now()

func RefreshMirrorEptToolchain(async bool) {
	lastRefreshTimeEpt = time.Now()
	closure := func() {
		log.Println("Start refreshing mirror ept toolchain")
		r, err := generateMirrorEptToolchain()
		if err != nil {
			log.Println("Failed to generate mirror ept toolchain", err)
		}
		mirrorEptToolchainCache = r
		log.Println("Refreshed mirror ept toolchain")
	}

	if async {
		go closure()
	} else {
		closure()
	}
}

func GetMirrorEptToolchain() vo.MirrorEptToolchain {
	// 如果超过 1 分钟则刷新
	if time.Since(lastRefreshTimeEpt) > 1*time.Minute {
		RefreshMirrorEptToolchain(true)
	}
	return mirrorEptToolchainCache
}
