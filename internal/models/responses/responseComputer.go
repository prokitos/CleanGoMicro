package responses

import (
	"modules/internal/models"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

// ответы для таблицы Computer

type ResponseComputer struct {
	Description string            `json:"description"       		 example:"description"`
	Code        int               `json:"code"               		 example:"status"`
	Computers   []tables.Computer `json:"computers,omitempty"       example:"...."`
}

func (instance ResponseComputer) GoodCreate() models.Response {
	return ResponseBase{}.GoodCreate("computer")
}
func (instance ResponseComputer) BadCreate() models.Response {
	return ResponseBase{}.BadCreate("computer")
}
func (instance ResponseComputer) GoodUpdate() models.Response {
	return ResponseBase{}.GoodUpdate("computer")
}
func (instance ResponseComputer) BadUpdate() models.Response {
	return ResponseBase{}.BadUpdate("computer")
}
func (instance ResponseComputer) GoodDelete() models.Response {
	return ResponseBase{}.GoodDelete("computer")
}
func (instance ResponseComputer) BadDelete() models.Response {
	return ResponseBase{}.BadDelete("computer")
}
func (instance ResponseComputer) GoodShow(curComputer []tables.Computer) models.Response {
	var items []models.Table
	for i := 0; i < len(curComputer); i++ {
		items = append(items, &curComputer[i])
	}
	return ResponseBase{}.GoodShow(items, "computer")
}
func (instance ResponseComputer) BadShow() models.Response {
	return ResponseBase{}.BadShow("computer")
}
func (instance ResponseComputer) InternalError() models.Response {
	return ResponseBase{}.InternalError()
}

func (instance ResponseComputer) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
func (instance ResponseComputer) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}
