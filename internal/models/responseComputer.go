package models

import "github.com/gofiber/fiber/v2"

type ResponseComputer struct {
	Description string     `json:"description"       		 example:"description"`
	Code        int        `json:"code"               		 example:"status"`
	Computers   []Computer `json:"computers,omitempty"       example:"...."`
}

func (instance ResponseComputer) GoodCreate() Response {
	instance.Code = 200
	instance.Description = "computer create success"
	instance.Computers = nil
	return instance.GetResponse()
}
func (instance ResponseComputer) BadCreate() Response {
	instance.Code = 400
	instance.Description = "computer create error"
	instance.Computers = nil
	return instance.GetResponse()
}
func (instance ResponseComputer) GoodUpdate() Response {
	instance.Code = 200
	instance.Description = "computer update success"
	instance.Computers = nil
	return instance.GetResponse()
}
func (instance ResponseComputer) BadUpdate() Response {
	instance.Code = 400
	instance.Description = "computer update error"
	instance.Computers = nil
	return instance.GetResponse()
}
func (instance ResponseComputer) GoodDelete() Response {
	instance.Code = 200
	instance.Description = "computer delete success"
	instance.Computers = nil
	return instance.GetResponse()
}
func (instance ResponseComputer) BadDelete() Response {
	instance.Code = 400
	instance.Description = "computer delete error"
	instance.Computers = nil
	return instance.GetResponse()
}
func (instance ResponseComputer) GoodShow(curComputer []Computer) Response {
	instance.Code = 200
	instance.Description = "computer show success"
	instance.Computers = curComputer
	return instance.GetResponse()
}
func (instance ResponseComputer) BadShow() Response {
	instance.Code = 400
	instance.Description = "computer show error"
	instance.Computers = nil
	return instance.GetResponse()
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

func (instance ResponseComputer) GetResponse() Response {
	var temp Response
	temp = instance
	return temp
}
