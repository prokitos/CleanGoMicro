package responses

import (
	"modules/internal/models"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

// ответы для таблицы Car

type ResponseCar struct {
	Description string       `json:"description"       		 example:"description"`
	Code        int          `json:"code"               		 example:"status"`
	Cars        []tables.Car `json:"cars,omitempty"       example:"...."`
}

func (instance ResponseCar) GoodCreate() models.Response {
	return ResponseBase{}.GoodCreate("car")
}
func (instance ResponseCar) BadCreate() models.Response {
	return ResponseBase{}.BadCreate("car")
}
func (instance ResponseCar) GoodUpdate() models.Response {
	return ResponseBase{}.GoodUpdate("car")
}
func (instance ResponseCar) BadUpdate() models.Response {
	return ResponseBase{}.BadUpdate("car")
}
func (instance ResponseCar) GoodDelete() models.Response {
	return ResponseBase{}.GoodDelete("car")
}
func (instance ResponseCar) BadDelete() models.Response {
	return ResponseBase{}.BadDelete("car")
}
func (instance ResponseCar) GoodShow(curCar []tables.Car) models.Response {
	var items []models.Table
	for i := 0; i < len(curCar); i++ {
		items = append(items, &curCar[i])
	}
	return ResponseBase{}.GoodShow(items, "car")
}
func (instance ResponseCar) BadShow() models.Response {
	return ResponseBase{}.BadShow("car")
}
func (instance ResponseCar) InternalError() models.Response {
	return ResponseBase{}.InternalError()
}

func (instance ResponseCar) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
func (instance ResponseCar) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}
