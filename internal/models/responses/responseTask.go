package responses

import (
	"modules/internal/models"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

type ResponseTask struct {
	Description string        `json:"description"       		 example:"description"`
	Code        int           `json:"code"               		 example:"status"`
	Tasks       []tables.Task `json:"tasks,omitempty"       example:"...."`
}

func (instance ResponseTask) GoodCreate() models.Response {
	return ResponseBase{}.GoodCreate("task")
}
func (instance ResponseTask) BadCreate() models.Response {
	return ResponseBase{}.BadCreate("task")
}
func (instance ResponseTask) GoodUpdate() models.Response {
	return ResponseBase{}.GoodUpdate("task")
}
func (instance ResponseTask) BadUpdate() models.Response {
	return ResponseBase{}.BadUpdate("task")
}
func (instance ResponseTask) GoodDelete() models.Response {
	return ResponseBase{}.GoodDelete("task")
}
func (instance ResponseTask) BadDelete() models.Response {
	return ResponseBase{}.BadDelete("task")
}
func (instance ResponseTask) GoodShow(curTask []tables.Task) models.Response {
	var items []models.Table
	for i := 0; i < len(curTask); i++ {
		items = append(items, &curTask[i])
	}
	return ResponseBase{}.GoodShow(items, "task")
}
func (instance ResponseTask) BadShow() models.Response {
	return ResponseBase{}.BadShow("task")
}
func (instance ResponseTask) InternalError() models.Response {
	return ResponseBase{}.InternalError()
}

func (instance ResponseTask) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
func (instance ResponseTask) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}
