package services

import (
	"module/internal/database"
	"module/internal/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

func GradeCreate(c *fiber.Ctx) error {

	var body models.Grades
	if err := c.BodyParser(&body); err != nil {
		log.Error("parse error")
		return models.ResponseBadRequest()
	}

	// сначала отправляем в бд. если чтото пошло не так, то в кафку не отправляем.
	err := database.CreateNewGrade(body)
	if err.Error() != models.ResponseGood().Error() {
		log.Error("database write error")
		return err
	}

	return producerSend(body)

}
