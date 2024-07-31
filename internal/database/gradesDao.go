package database

import (
	"module/internal/models"

	log "github.com/sirupsen/logrus"
)

func CreateNewGrade(curUser models.Grades) error {

	if result := GlobalHandler.Create(&curUser); result.Error != nil {
		log.Debug("create record error!")
		return models.ResponseErrorAtServer()
	}

	return models.ResponseGood()
}
