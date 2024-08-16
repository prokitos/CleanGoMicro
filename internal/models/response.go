package models

import (
	"github.com/gofiber/fiber/v2"
)

type ResponseUser struct {
	Description string `json:"description"        example:"description"`
	Code        int    `json:"code"               example:"status"`
	Users       []User `json:"users,omitempty"    example:"...."`
}

type ResponseComputer struct {
	Description string     `json:"description"       		 example:"description"`
	Code        int        `json:"code"               		 example:"status"`
	Computers   []Computer `json:"computers,omitempty"       example:"...."`
}

func (instance *ResponseUser) GoodCreate() {
	instance.Code = 200
	instance.Description = "user create success"
	instance.Users = nil
}
func (instance *ResponseUser) BadCreate() {
	instance.Code = 400
	instance.Description = "user create error"
	instance.Users = nil
}
func (instance *ResponseUser) GoodShow(curUser []User) {
	instance.Code = 200
	instance.Description = "user show success"
	instance.Users = curUser
}
func (instance *ResponseUser) BadShow() {
	instance.Code = 400
	instance.Description = "user show error"
	instance.Users = nil
}

func (instance *ResponseUser) GetResponse() ResponseUser {
	return *instance
}
func (instance *ResponseUser) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}

func (instance *ResponseComputer) GoodCreate() {
	instance.Code = 200
	instance.Description = "computer create success"
	instance.Computers = nil
}
func (instance *ResponseComputer) BadCreate() {
	instance.Code = 400
	instance.Description = "computer create error"
	instance.Computers = nil
}
func (instance *ResponseComputer) GoodShow(curComputer []Computer) {
	instance.Code = 200
	instance.Description = "computer show success"
	instance.Computers = curComputer
}
func (instance *ResponseComputer) BadShow() {
	instance.Code = 400
	instance.Description = "computer show error"
	instance.Computers = nil
}

func (instance *ResponseComputer) GetResponse() ResponseComputer {
	return *instance
}

func (instance *ResponseComputer) GetError(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}
