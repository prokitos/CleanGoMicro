package transport

import (
	"modules/internal/models/tables"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

func getCar(c *fiber.Ctx) error {
	var curCar tables.Car
	curCar.GetQueryParams(c)
	curCar.GetQueryId(c)
	return services.CarShow(curCar).GetError(c)
}

func insertCar(c *fiber.Ctx) error {
	var curCar tables.Car
	curCar.GetQueryParams(c)
	return services.CarInsert(curCar).GetError(c)
}

func deleteCar(c *fiber.Ctx) error {
	var curCar tables.Car
	curCar.GetQueryId(c)
	return services.CarDelete(curCar).GetError(c)
}

func updateCar(c *fiber.Ctx) error {
	var curCar tables.Car
	curCar.GetQueryParams(c)
	curCar.GetQueryId(c)
	return services.CarUpdate(curCar).GetError(c)
}
