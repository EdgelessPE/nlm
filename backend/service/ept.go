package service

import (
	"fmt"
	"nlm/db"
	"nlm/model"
	"nlm/vo"

	"github.com/stoewer/go-strcase"
)

func GetEpts(params vo.GetEptsParams) ([]model.Ept, int64, error) {
	var epts []model.Ept
	var total int64

	tx := db.DB.Model(&model.Ept{})

	if params.Q != "" {
		tx = tx.Where("LOWER(name) LIKE LOWER(?)", "%"+params.Q+"%")
	}

	if params.Sort != 0 {
		var order string
		if params.Sort == 1 {
			order = "ASC"
		} else {
			order = "DESC"
		}
		tx = tx.Order(fmt.Sprintf("%s %s", strcase.SnakeCase(params.SortBy), order))
	}

	tx.Count(&total)
	tx.Offset(params.Offset).Limit(params.Limit).Find(&epts)

	return epts, total, nil
}
