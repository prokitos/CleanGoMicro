package responses

import (
	"modules/internal/models"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2"
)

type ResponseUser struct {
	Description string        `json:"description"        example:"description"`
	Code        int           `json:"code"               example:"status"`
	Users       []tables.User `json:"users,omitempty"    example:"...."`
}

func (instance ResponseUser) GoodCreate() models.Response {
	instance.Code = 200
	instance.Description = "user create success"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) BadCreate() models.Response {
	instance.Code = 400
	instance.Description = "user create error"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) GoodUpdate() models.Response {
	instance.Code = 200
	instance.Description = "user update success"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) BadUpdate() models.Response {
	instance.Code = 400
	instance.Description = "user update error"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) GoodDelete() models.Response {
	instance.Code = 200
	instance.Description = "user delete success"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) BadDelete() models.Response {
	instance.Code = 400
	instance.Description = "user delete error"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) GoodShow(curUser []tables.User) models.Response {
	instance.Code = 200
	instance.Description = "user show success"
	instance.Users = curUser
	return instance.GetResponse()
}
func (instance ResponseUser) BadShow() models.Response {
	instance.Code = 400
	instance.Description = "user show error"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) InternalError() models.Response {
	instance.Code = 400
	instance.Description = "internal error"
	instance.Users = nil
	return instance.GetResponse()
}

func (instance ResponseUser) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}

func (instance ResponseUser) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}

func (instance ResponseUser) GetResponse() models.Response {
	var temp models.Response
	temp = instance
	return temp
}
