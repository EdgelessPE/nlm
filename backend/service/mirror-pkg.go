package service

import (
	"nlm/config"
	"nlm/constant"
	"nlm/vo"
	"time"
)

func GenerateMirrorPkgSoftware() (vo.MirrorPkgSoftware, error) {
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
			r = append(r, vo.MirrorPkgSoftwareRelease{
				FileName:  release.FileName,
				Size:      release.FileSize,
				Timestamp: release.CreatedAt.UnixMilli(),
				Version:   release.Version,
				Meta:      release.Meta,
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
