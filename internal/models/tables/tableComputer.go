package tables

import (
	"modules/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Computer struct {
	Computer_id string `json:"computer_id,omitempty" example:"12" bson:"_id,omitempty"`
	Price       int    `json:"price,omitempty" example:"125000" bson:"price,omitempty"`
	Ram         string `json:"ram,omitempty" example:"16gb" bson:"ram,omitempty"`
	Cpu         string `json:"cpu,omitempty" example:"intel core i7" bson:"cpu,omitempty"`
	Gpu         string `json:"gpu,omitempty" example:"geforce gtx 1080" bson:"gpu,omitempty"`
}

type ComputerOutput struct {
	Computer_id primitive.ObjectID `json:"computer_id" bson:"_id,omitempty"`
	Price       int                `json:"price,omitempty" example:"125000" bson:"price,omitempty"`
	Ram         string             `json:"ram,omitempty" example:"16gb" bson:"ram,omitempty"`
	Cpu         string             `json:"cpu,omitempty" example:"intel core i7" bson:"cpu,omitempty"`
	Gpu         string             `json:"gpu,omitempty" example:"geforce gtx 1080" bson:"gpu,omitempty"`
}

func (instance *Computer) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *Computer) RecordShow(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	err := dao.ShowData(instance, db)
	return err
}
func (instance *Computer) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *Computer) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
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

func (instance *Computer) OutputGet() ComputerOutput {
	var newComputer ComputerOutput
	newComputer.Cpu = instance.Cpu
	newComputer.Gpu = instance.Gpu
	newComputer.Price = instance.Price
	newComputer.Ram = instance.Ram

	if instance.Computer_id != "" {
		objID, err := primitive.ObjectIDFromHex(instance.Computer_id)
		if err != nil {
			return newComputer
		}

		newComputer.Computer_id = objID
	}

	return newComputer
}
