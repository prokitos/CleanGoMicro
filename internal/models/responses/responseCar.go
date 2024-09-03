package responses

import (
	"modules/internal/models"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

type ResponseCar struct {
	Description string       `json:"description"       		 example:"description"`
	Code        int          `json:"code"               		 example:"status"`
	Cars        []tables.Car `json:"cars,omitempty"       example:"...."`
}

func (instance ResponseCar) GoodCreate() models.Response {
	instance.Code = 200
	instance.Description = "car create success"
	instance.Cars = nil
	return instance.GetResponse()
}
func (instance ResponseCar) BadCreate() models.Response {
	instance.Code = 400
	instance.Description = "car create error"
	instance.Cars = nil
	return instance.GetResponse()
}
func (instance ResponseCar) GoodUpdate() models.Response {
	instance.Code = 200
	instance.Description = "car update success"
	instance.Cars = nil
	return instance.GetResponse()
}
func (instance ResponseCar) BadUpdate() models.Response {
	instance.Code = 400
	instance.Description = "car update error"
	instance.Cars = nil
	return instance.GetResponse()
}
func (instance ResponseCar) GoodDelete() models.Response {
	instance.Code = 200
	instance.Description = "car delete success"
	instance.Cars = nil
	return instance.GetResponse()
}
func (instance ResponseCar) BadDelete() models.Response {
	instance.Code = 400
	instance.Description = "car delete error"
	instance.Cars = nil
	return instance.GetResponse()
}
func (instance ResponseCar) GoodShow(curCar []tables.Car) models.Response {
	instance.Code = 200
	instance.Description = "car show success"
	instance.Cars = curCar
	return instance.GetResponse()
}
func (instance ResponseCar) BadShow() models.Response {
	instance.Code = 400
	instance.Description = "car show error"
	instance.Cars = nil
	return instance.GetResponse()
}
func (instance ResponseCar) InternalError() models.Response {
	instance.Code = 400
	instance.Description = "internal error"
	instance.Cars = nil
	return instance.GetResponse()
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

func (instance ResponseCar) GetResponse() models.Response {
	var temp models.Response
	temp = instance
	return temp
}
