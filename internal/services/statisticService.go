package services

import (
	"module/internal/database"

	"github.com/gofiber/fiber/v2"
)

func StatisticGet(c *fiber.Ctx) error {

	res, err := database.GetStatistics()

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"status": err.Error(), "data": res})

}
