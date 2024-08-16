package transport

import (
	"modules/internal/models"

	"github.com/gofiber/fiber/v2"
)

func getUser(c *fiber.Ctx) error {
	var newResponse models.ResponseUser
	newResponse.GoodShow([]models.User{{Login: "vasya", Password: "123456"}, {Login: "zhora", Password: "2238"}})
	return newResponse.GetError(c)
}

func insertUser(c *fiber.Ctx) error {
	var newResponse models.ResponseUser
	newResponse.GoodCreate()
	return newResponse.GetError(c)
}
