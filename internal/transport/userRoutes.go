package transport

import (
	"modules/internal/models"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

func getUser(c *fiber.Ctx) error {
	var curUser models.User
	curUser.GetQueryParams(c)
	curUser.GetQueryId(c)
	return services.UserShow(curUser).GetError(c)
}

func insertUser(c *fiber.Ctx) error {
	var curUser models.User
	curUser.GetQueryParams(c)
	return services.UserInsert(curUser).GetError(c)
}

func deleteUser(c *fiber.Ctx) error {
	var curUser models.User
	curUser.GetQueryId(c)
	return services.UserDelete(curUser).GetError(c)
}

func updateUser(c *fiber.Ctx) error {
	var curUser models.User
	curUser.GetQueryParams(c)
	curUser.GetQueryId(c)
	return services.UserUpdate(curUser).GetError(c)
}
