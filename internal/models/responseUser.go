package models

import "github.com/gofiber/fiber/v2"

type ResponseUser struct {
	Description string  `json:"description"        example:"description"`
	Code        int     `json:"code"               example:"status"`
	Users       []Table `json:"users,omitempty"    example:"...."`
}

func (instance ResponseUser) GoodCreate() Response {
	instance.Code = 200
	instance.Description = "user create success"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) BadCreate() Response {
	instance.Code = 400
	instance.Description = "user create error"
	instance.Users = nil
	return instance.GetResponse()
}
func (instance ResponseUser) GoodShow(curUser []Table) Response {
	instance.Code = 200
	instance.Description = "user show success"
	instance.Users = curUser
	return instance.GetResponse()
}
func (instance ResponseUser) BadShow() Response {
	instance.Code = 400
	instance.Description = "user show error"
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

func (instance ResponseUser) GetResponse() Response {
	var temp Response
	temp = instance
	return temp
}
