package transport

import (
	"modules/internal/models"

	"github.com/gofiber/fiber/v2"
)

func getUser(c *fiber.Ctx) error {
	var newResponse models.ResponseUser

	var builder models.TableBuilder
	var temp []models.Table
	temp = append(temp, builder.UserCreate("vanya", "123456").Instance)
	temp = append(temp, builder.UserCreate("seg", "wey").Instance)

	newResponse.GoodShow(temp)
	return newResponse.GetError(c)
}

func insertUser(c *fiber.Ctx) error {
	var newResponse models.ResponseUser
	newResponse.GoodCreate()
	return newResponse.GetError(c)
}
