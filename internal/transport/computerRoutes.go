package transport

import (
	"modules/internal/models"

	"github.com/gofiber/fiber/v2"
)

func getComputer(c *fiber.Ctx) error {
	var newResponse models.ResponseComputer
	newResponse.BadShow()
	return newResponse.GetError(c)
}

func insertComputer(c *fiber.Ctx) error {
	var newResponse models.ResponseComputer
	newResponse.GoodCreate()
	return newResponse.GetError(c)
}
