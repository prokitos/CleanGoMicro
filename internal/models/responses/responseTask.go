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
	instance.Code = 200
	instance.Description = "task create success"
	instance.Tasks = nil
	return instance.GetResponse()
}
func (instance ResponseTask) BadCreate() models.Response {
	instance.Code = 400
	instance.Description = "task create error"
	instance.Tasks = nil
	return instance.GetResponse()
}
func (instance ResponseTask) GoodUpdate() models.Response {
	instance.Code = 200
	instance.Description = "task update success"
	instance.Tasks = nil
	return instance.GetResponse()
}
func (instance ResponseTask) BadUpdate() models.Response {
	instance.Code = 400
	instance.Description = "task update error"
	instance.Tasks = nil
	return instance.GetResponse()
}
func (instance ResponseTask) GoodDelete() models.Response {
	instance.Code = 200
	instance.Description = "task delete success"
	instance.Tasks = nil
	return instance.GetResponse()
}
func (instance ResponseTask) BadDelete() models.Response {
	instance.Code = 400
	instance.Description = "task delete error"
	instance.Tasks = nil
	return instance.GetResponse()
}
func (instance ResponseTask) GoodShow(curTask []tables.Task) models.Response {
	instance.Code = 200
	instance.Description = "task show success"
	instance.Tasks = curTask
	return instance.GetResponse()
}
func (instance ResponseTask) BadShow() models.Response {
	instance.Code = 400
	instance.Description = "task show error"
	instance.Tasks = nil
	return instance.GetResponse()
}
func (instance ResponseTask) InternalError() models.Response {
	instance.Code = 400
	instance.Description = "internal error"
	instance.Tasks = nil
	return instance.GetResponse()
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

func (instance ResponseTask) GetResponse() models.Response {
	var temp models.Response
	temp = instance
	return temp
}
