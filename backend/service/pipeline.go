package service

import (
	"fmt"
	"nlm/db"
	"nlm/model"
	"nlm/vo"

	"github.com/stoewer/go-strcase"
)

func GetPipelines(params vo.PipelineParams) ([]model.Pipeline, int64, error) {
	var pipelines []model.Pipeline
	var total int64

	tx := db.DB.Model(&model.Pipeline{})

	if params.ModelName != "" {
		tx = tx.Where("model_name = ?", params.ModelName)
	}
	if params.Status != "" {
		tx = tx.Where("status = ?", params.Status)
	}

	if params.Sort != 0 {
		var order string
		if params.Sort == 1 {
			order = "ASC"
		} else {
			order = "DESC"
		}
		tx = tx.Order(fmt.Sprintf("%s %s", strcase.SnakeCase(params.SortBy), order))
	} else {
		// 默认按照创建时间降序排序
		tx = tx.Order("created_at DESC")
	}

	tx.Count(&total)
	tx.Offset(params.Offset).Limit(params.Limit).Find(&pipelines)

	return pipelines, total, nil
}
