package service

import (
	"fmt"
	"nlm/db"
	"nlm/model"
)

func AddStorage(sourceFilePath string) (string, error) {
	var s model.Storage
	s.SourceFilePath = sourceFilePath
	db.DB.Create(&s)
	return fmt.Sprint(s.ID), nil
}
