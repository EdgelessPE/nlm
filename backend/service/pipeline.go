package service

import (
	"nlm/db"
	"nlm/model"
	"nlm/vo"
)

func GetPipelines(params vo.PipelineParams) ([]model.Pipeline, int64, error) {
	var pipelines []model.Pipeline
	var total int64

	tx := db.DB.Model(&model.Pipeline{})

	if params.ModelName != "" {
		tx = tx.Where("model_name = ?", params.ModelName)
	}

	tx.Count(&total)
	tx.Offset(params.Offset).Limit(params.Limit).Find(&pipelines)

	return pipelines, total, nil
}
