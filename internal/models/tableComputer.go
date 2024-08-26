package models

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ComputerBase struct {
	Price int    `json:"price,omitempty" example:"125000"`
	Ram   string `json:"ram,omitempty" example:"16gb"`
	Cpu   string `json:"cpu,omitempty" example:"intel core i7"`
	Gpu   string `json:"gpu,omitempty" example:"geforce gtx 1080"`
}

type Computer struct {
	Computer_id string `json:"computer_id,omitempty" example:"12" bson:"_id,omitempty"`
	ComputerBase
}

type ComputerOutput struct {
	Computer_id primitive.ObjectID `json:"computer_id" bson:"_id,omitempty"`
	ComputerBase
}

func (instance *Computer) RecordCreate(db GormDatabase) Response {
	return db.CreateData(instance)
}
func (instance *Computer) RecordShow(db GormDatabase) Response {
	err := db.ShowData(instance)
	return err
}
func (instance *Computer) RecordDelete(db GormDatabase) Response {
	return db.DeleteData(instance)
}
func (instance *Computer) RecordUpdate(db GormDatabase) Response {
	return db.UpdateData(instance)
}
func (instance *Computer) GetId() string {
	return instance.Computer_id
}

func (instance *Computer) GetQueryId(c *fiber.Ctx) error {
	id := c.Query("id", "")
	instance.Computer_id = id
	return nil
}

func (instance *Computer) GetQueryParams(c *fiber.Ctx) error {
	price, err := strconv.Atoi(c.Query("price", ""))
	if err != nil && price != 0 {
		return err
	}
	instance.Price = price
	instance.Ram = c.Query("ram", "")
	instance.Cpu = c.Query("cpu", "")
	instance.Gpu = c.Query("gpu", "")
	return nil
}

func (instance *Computer) GetBodyParams(c *fiber.Ctx) error {
	if err := c.BodyParser(&instance); err != nil {
		return err
	}
	return nil
}
