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

func GetNeps() ([]model.Nep, error) {
	var neps []model.Nep
	db.DB.Find(&neps)
	return neps, nil
}

func GetNepsWithPagination(offset int, limit int) ([]model.Nep, error) {
	var neps []model.Nep
	db.DB.Offset(offset).Limit(limit).Find(&neps)
	return neps, nil
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

func GetRelease(scope string, name string, fileName string) (model.Release, error) {
	n, err := GetNep(scope, name)
	if err != nil {
		return model.Release{}, err
	}
	var release model.Release
	db.DB.Where("nep_id = ? AND file_name = ?", n.ID.String(), fileName).First(&release)
	if release.ID == uuid.Nil {
		return release, errors.New("release not found")
	}
	return release, nil
}
