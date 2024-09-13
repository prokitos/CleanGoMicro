package tables

import (
	"modules/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// таблица Product. Методы для получения REST данных, а также выполнение команд в нужном DAO. Вызывается из сервисов.

type Product struct {
	Product_id int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Title      string `json:"title" example:"something"`
	Price      int    `json:"price" example:"1400"`
}

func (instance *Product) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *Product) RecordShow(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	err := dao.ShowData(instance, db)
	return err
}
func (instance *Product) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *Product) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
}
func (instance *Product) GetId() int {
	return instance.Product_id
}

func (instance *Product) GetQueryId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id", ""))
	if err != nil {
		return err
	}
	instance.Product_id = id
	return nil
}

func (instance *Product) GetQueryParams(c *fiber.Ctx) error {
	instance.Title = c.Query("title", "")
	price, err := strconv.Atoi(c.Query("price", ""))
	if err != nil && price != 0 {
		return err
	}
	instance.Price = price
	return nil
}

func (instance *Product) GetBodyParams(c *fiber.Ctx) error {
	if err := c.BodyParser(&instance); err != nil {
		return err
	}
	return nil
}
