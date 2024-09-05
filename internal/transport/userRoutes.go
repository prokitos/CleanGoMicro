package transport

import (
	"modules/internal/models/tables"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

// роуты для users

func getUser(c *fiber.Ctx) error {
	var curUser tables.User
	curUser.GetQueryParams(c)
	curUser.GetQueryId(c)
	return services.UserShow(curUser).GetError(c)
}

func insertUser(c *fiber.Ctx) error {
	var curUser tables.User
	curUser.GetQueryParams(c)
	return services.UserInsert(curUser).GetError(c)
}

func deleteUser(c *fiber.Ctx) error {
	var curUser tables.User
	curUser.GetQueryId(c)
	return services.UserDelete(curUser).GetError(c)
}

func updateUser(c *fiber.Ctx) error {
	var curUser tables.User
	curUser.GetQueryParams(c)
	curUser.GetQueryId(c)
	return services.UserUpdate(curUser).GetError(c)
}
