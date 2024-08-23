package transport

import (
	"fmt"
	"modules/internal/models"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

func getComputer(c *fiber.Ctx) error {
	var curComputer models.Computer
	curComputer.GetQueryParams(c)
	curComputer.GetQueryId(c)
	fmt.Println("zdes")
	fmt.Println(curComputer)
	return services.ComputerShow(curComputer).GetError(c)
}

func insertComputer(c *fiber.Ctx) error {
	var curComputer models.Computer
	curComputer.GetQueryParams(c)
	return services.ComputerInsert(curComputer).GetError(c)
}

func deleteComputer(c *fiber.Ctx) error {
	var curComputer models.Computer
	curComputer.GetQueryId(c)
	return services.ComputerDelete(curComputer).GetError(c)
}

func updateComputer(c *fiber.Ctx) error {
	var curComputer models.Computer
	curComputer.GetQueryParams(c)
	curComputer.GetQueryId(c)
	return services.ComputerUpdate(curComputer).GetError(c)
}
