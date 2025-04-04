package service

import (
	"errors"
	"nlm/db"
	"nlm/model"

	"github.com/google/uuid"
)

func HasNep(scope string, name string) bool {
	var nep model.Nep
	db.DB.Where("scope = ? AND name = ?", scope, name).First(&nep)
	return nep.ID != uuid.Nil
}

func AddNep(scope string, name string) (model.Nep, error) {
	if HasNep(scope, name) {
		return model.Nep{}, errors.New("nep already exists")
	}
	r := model.Nep{Scope: scope, Name: name}
	db.DB.Create(&r)
	return r, nil
}

func GetNep(scope string, name string) (model.Nep, error) {
	var nep model.Nep
	db.DB.Where("scope = ? AND name = ?", scope, name).First(&nep)
	if nep.ID == uuid.Nil {
		return nep, errors.New("nep not found")
	}
	return nep, nil
}

func GetNeps(offset int, limit int) ([]model.Nep, error) {
	var neps []model.Nep
	db.DB.Offset(offset).Limit(limit).Find(&neps)
	return neps, nil
}

// 默认添加的 release 版本号是最高的
func AddRelease(scope string, name string, version string, flags string, fileName string, pipelineId string) (model.Release, error) {
	n, err := GetNep(scope, name)
	if err != nil {
		return model.Release{}, err
	}
	r := model.Release{Version: version, Flags: flags, FileName: fileName, PipelineId: pipelineId, NepId: n.ID.String()}
	db.DB.Create(&r)
	n.LatestReleaseVersion = version
	db.DB.Save(&n)
	return r, nil
}

func GetReleases(scope string, name string) ([]model.Release, error) {
	// 获取 Nep
	n, err := GetNep(scope, name)
	if err != nil {
		return nil, err
	}

	// 获取 Releases
	var releases []model.Release
	db.DB.Where("nep_id = ?", n.ID.String()).Find(&releases)
	return releases, nil
}
