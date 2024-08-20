package models

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type Computer struct {
	Computer_id int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Price       int    `json:"price" example:"125000"`
	Ram         string `json:"ram" example:"16gb"`
	Cpu         string `json:"cpu" example:"intel core i7"`
	Gpu         string `json:"gpu" example:"geforce gtx 1080"`
}

func (instance *Computer) RecordCreate(db GormDatabase) Response {
	return db.CreateData(instance)
}
func (instance *Computer) RecordShow(db GormDatabase) Response {
	_, err := db.ShowData(instance)
	return err
}
func (instance *Computer) RecordDelete(db GormDatabase) Response {
	return db.DeleteData(instance)
}
func (instance *Computer) RecordUpdate(db GormDatabase) Response {
	return db.UpdateData(instance)
}

func (instance *Computer) GetQueryId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id", ""))
	if err != nil {
		return err
	}
	instance.Computer_id = id
	return nil
}

func (instance *Computer) GetQueryParams(c *fiber.Ctx) error {
	price, err := strconv.Atoi(c.Query("price", ""))
	if err != nil {
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
