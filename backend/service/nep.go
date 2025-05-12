package service

import (
	"errors"
	"log"
	"nlm/db"
	"nlm/model"
	"nlm/vo"
	"time"

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

func GetNeps(params vo.NepParams) ([]model.Nep, int64, error) {
	var neps []model.Nep
	var total int64

	tx := db.DB.Model(&model.Nep{})
	if params.Q != "" {
		tx = tx.Where("scope LIKE ? OR name LIKE ?", "%"+params.Q+"%", "%"+params.Q+"%")
	}
	if params.Scope != "" {
		tx = tx.Where("scope = ?", params.Scope)
	}

	tx.Count(&total)
	if params.Offset >= 0 && params.Limit > 0 {
		tx = tx.Offset(params.Offset).Limit(params.Limit)
	}
	tx.Find(&neps)

	return neps, total, nil
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

func GetReleases(params vo.ReleaseParams) ([]model.Release, int64, error) {
	var releases []model.Release
	var total int64

	tx := db.DB.Model(&model.Release{})
	if params.Q != "" {
		tx = tx.Where("file_name LIKE ?", "%"+params.Q+"%")
	}
	if params.NepID != "" {
		tx = tx.Where("nep_id = ?", params.NepID)
	}
	if params.IsBotSuccess {
		tx = tx.Where("is_bot_success = true")
	}
	if params.IsQaSuccess {
		tx = tx.Where("is_qa_success = true")
	}
	if params.Version != "" {
		tx = tx.Where("version = ?", params.Version)
	}

	tx.Count(&total)
	if params.Offset >= 0 && params.Limit > 0 {
		tx = tx.Offset(params.Offset).Limit(params.Limit)
	}
	tx.Find(&releases)

	return releases, total, nil
}

func CleanOutdatedRelease() error {
	log.Println("Cleaning outdated release..")
	// 删除更新时间大于 30 天且不是最后一个大版本的 Release
	var releases []model.Release
	db.DB.Where("updated_at < ? AND is_last_major = false", time.Now().AddDate(0, 0, -30)).Find(&releases)
	for _, release := range releases {
		log.Printf("Cleaning outdated release: %s (%s)", release.ID.String(), release.FileName)
		db.DB.Delete(&release)
		DeleteStorage(release.StorageKey)
	}
	log.Println("Cleaned outdated release")
	return nil
}
