package tables

import (
	"modules/internal/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// таблица Car. Методы для получения REST данных, а также выполнение команд в нужном DAO. Вызывается из сервисов.

type Car struct {
	Car_id string `json:"car_id,omitempty" example:"12" bson:"_id,omitempty"`
	Mark   string `json:"mark" example:"lada" bson:"mark,omitempty"`
	Year   string `json:"year" example:"1990" bson:"year,omitempty"`
}
type CarOutput struct {
	Car_id primitive.ObjectID `json:"car_id" bson:"_id,omitempty"`
	Mark   string             `json:"mark" example:"lada" bson:"mark,omitempty"`
	Year   string             `json:"year" example:"1990" bson:"year,omitempty"`
}

func (instance *Car) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *Car) RecordShow(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	err := dao.ShowData(instance, db)
	return err
}
func (instance *Car) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *Car) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
}
func (instance *Car) GetId() string {
	return instance.Car_id
}

func (instance *Car) GetQueryId(c *fiber.Ctx) error {
	id := c.Query("id", "")
	instance.Car_id = id
	return nil
}

func (instance *Car) GetQueryParams(c *fiber.Ctx) error {
	instance.Mark = c.Query("mark", "")
	instance.Year = c.Query("year", "")
	return nil
}

func (instance *Car) GetBodyParams(c *fiber.Ctx) error {
	if err := c.BodyParser(&instance); err != nil {
		return err
	}
	return nil
}

func (instance *Car) OutputGet() CarOutput {
	var newCar CarOutput
	newCar.Mark = instance.Mark
	newCar.Year = instance.Year

	if instance.Car_id != "" {
		objID, err := primitive.ObjectIDFromHex(instance.Car_id)
		if err != nil {
			return newCar
		}

		newCar.Car_id = objID
	}

	return newCar
}
