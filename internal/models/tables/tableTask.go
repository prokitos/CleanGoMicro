package tables

import (
	"modules/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// таблица Task. Методы для получения REST данных, а также выполнение команд в нужном DAO. Вызывается из сервисов.

type Task struct {
	Task_id     int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Description string `json:"description" example:"something"`
	DueDate     string `json:"due_date" example:"01.01.2010"`
}

func (instance *Task) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *Task) RecordShow(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	err := dao.ShowData(instance, db)
	return err
}
func (instance *Task) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *Task) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
}
func (instance *Task) GetId() int {
	return instance.Task_id
}

func (instance *Task) GetQueryId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id", ""))
	if err != nil {
		return err
	}
	instance.Task_id = id
	return nil
}

func (instance *Task) GetQueryParams(c *fiber.Ctx) error {
	instance.Description = c.Query("description", "")
	instance.DueDate = c.Query("due_date", "")
	return nil
}

func (instance *Task) GetBodyParams(c *fiber.Ctx) error {
	if err := c.BodyParser(&instance); err != nil {
		return err
	}
	return nil
}
