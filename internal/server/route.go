package server

import (
	"module/internal/models"
	"module/internal/services"

	"github.com/gofiber/fiber/v2"
)

func routeAddTask(c *fiber.Ctx) error {

	return services.GradeCreate(c)

}

func routeGetStatistic(c *fiber.Ctx) error {

	//return services.StatisticGet(c)
	return models.ResponseGood()
}
