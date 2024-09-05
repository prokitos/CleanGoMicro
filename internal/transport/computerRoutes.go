package transport

import (
	"modules/internal/models/tables"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

// роуты для computers

func getComputer(c *fiber.Ctx) error {
	var curComputer tables.Computer
	curComputer.GetQueryParams(c)
	curComputer.GetQueryId(c)
	return services.ComputerShow(curComputer).GetError(c)
}

func insertComputer(c *fiber.Ctx) error {
	var curComputer tables.Computer
	curComputer.GetQueryParams(c)
	return services.ComputerInsert(curComputer).GetError(c)
}

func deleteComputer(c *fiber.Ctx) error {
	var curComputer tables.Computer
	curComputer.GetQueryId(c)
	return services.ComputerDelete(curComputer).GetError(c)
}

func updateComputer(c *fiber.Ctx) error {
	var curComputer tables.Computer
	curComputer.GetQueryParams(c)
	curComputer.GetQueryId(c)
	return services.ComputerUpdate(curComputer).GetError(c)
}
