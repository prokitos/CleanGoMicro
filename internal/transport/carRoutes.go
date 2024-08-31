package transport

import (
	"modules/internal/models/tables"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

func insertCar(c *fiber.Ctx) error {
	var curCar tables.Car
	curCar.Mark = "lada"
	curCar.Year = "2000"
	return services.CarInsert(curCar).GetError(c)
}
