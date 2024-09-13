package responses

import (
	"modules/internal/models"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

// ответы для таблицы Product

type ResponseProduct struct {
	Description string           `json:"description"       			 example:"description"`
	Code        int              `json:"code"               		 example:"status"`
	Product     []tables.Product `json:"tasks,omitempty"       		 example:"...."`
}

func (instance ResponseProduct) GoodCreate() models.Response {
	return ResponseBase{}.GoodCreate("product")
}
func (instance ResponseProduct) BadCreate() models.Response {
	return ResponseBase{}.BadCreate("product")
}
func (instance ResponseProduct) GoodUpdate() models.Response {
	return ResponseBase{}.GoodUpdate("product")
}
func (instance ResponseProduct) BadUpdate() models.Response {
	return ResponseBase{}.BadUpdate("product")
}
func (instance ResponseProduct) GoodDelete() models.Response {
	return ResponseBase{}.GoodDelete("product")
}
func (instance ResponseProduct) BadDelete() models.Response {
	return ResponseBase{}.BadDelete("product")
}
func (instance ResponseProduct) GoodShow(curProd []tables.Product) models.Response {
	var items []models.Table
	for i := 0; i < len(curProd); i++ {
		items = append(items, &curProd[i])
	}
	return ResponseBase{}.GoodShow(items, "product")
}
func (instance ResponseProduct) BadShow() models.Response {
	return ResponseBase{}.BadShow("product")
}
func (instance ResponseProduct) InternalError() models.Response {
	return ResponseBase{}.InternalError()
}

func (instance ResponseProduct) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
func (instance ResponseProduct) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}
