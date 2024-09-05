package tables

import (
	"modules/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// таблица User. Методы для получения REST данных, а также выполнение команд в нужном DAO. Вызывается из сервисов.

type User struct {
	User_id  int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Login    string `json:"login" example:"admin"`
	Password string `json:"password" example:"123456"`
}

func (instance *User) RecordCreate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.CreateData(instance, db)
}
func (instance *User) RecordShow(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	err := dao.ShowData(instance, db)
	return err
}
func (instance *User) RecordDelete(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.DeleteData(instance, db)
}
func (instance *User) RecordUpdate(db models.DatabaseCore, dao models.DatabaseDao) models.Response {
	return dao.UpdateData(instance, db)
}
func (instance *User) GetId() int {
	return instance.User_id
}

func (instance *User) GetQueryId(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Query("id", ""))
	if err != nil {
		return err
	}
	instance.User_id = id
	return nil
}

func (instance *User) GetQueryParams(c *fiber.Ctx) error {
	instance.Login = c.Query("login", "")
	instance.Password = c.Query("password", "")
	return nil
}

func (instance *User) GetBodyParams(c *fiber.Ctx) error {
	if err := c.BodyParser(&instance); err != nil {
		return err
	}
	return nil
}
