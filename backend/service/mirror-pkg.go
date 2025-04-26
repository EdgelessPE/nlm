package service

import (
	"encoding/json"
	"nlm/config"
	"nlm/constant"
	"nlm/vo"
	"time"
)

var mirrorPkgSoftwareCache vo.MirrorPkgSoftware

func generateMirrorPkgSoftware() (vo.MirrorPkgSoftware, error) {
	tree := make(map[string][]vo.MirrorPkgSoftwareTreeItem)

	neps, err := GetNeps()
	if err != nil {
		return vo.MirrorPkgSoftware{}, err
	}

	for _, nep := range neps {
		// 获取 releases
		releases, err := GetReleases(nep.Scope, nep.Name)
		if err != nil {
			return vo.MirrorPkgSoftware{}, err
		}

		// 转换 release 类型
		r := make([]vo.MirrorPkgSoftwareRelease, 0)
		for _, release := range releases {
			var meta interface{}
			err = json.Unmarshal(release.Meta, &meta)
			if err != nil {
				return vo.MirrorPkgSoftware{}, err
			}
			r = append(r, vo.MirrorPkgSoftwareRelease{
				FileName:  release.FileName,
				Size:      release.FileSize,
				Timestamp: release.CreatedAt.UnixMilli(),
				Version:   release.Version,
				Meta:      meta,
			})
		}

		// 添加到 tree
		tree[nep.Scope] = append(tree[nep.Scope], vo.MirrorPkgSoftwareTreeItem{
			Name:     nep.Name,
			Releases: r,
		})
	}

	return vo.MirrorPkgSoftware{
		Timestamp:   time.Now().UnixMilli(),
		URLTemplate: config.ENV.ROOT_URL + constant.API_PREFIX + constant.ServicePathRedirectTemplate,
		Tree:        tree,
	}, nil
}

var lastRefreshTime time.Time = time.Now()

func RefreshMirrorPkgSoftware(async bool) {
	lastRefreshTime = time.Now()
	closure := func() {
		println("Start refreshing mirror pkg software")
		r, err := generateMirrorPkgSoftware()
		if err != nil {
			println("Failed to generate mirror pkg software", err)
		}
		mirrorPkgSoftwareCache = r
		println("Refreshed mirror pkg software")
	}

	if async {
		go closure()
	} else {
		closure()
	}
}

func GetMirrorPkgSoftware() vo.MirrorPkgSoftware {
	// 如果超过 1 分钟则刷新
	if time.Since(lastRefreshTime) > 1*time.Minute {
		RefreshMirrorPkgSoftware(true)
	}
	return mirrorPkgSoftwareCache
}
