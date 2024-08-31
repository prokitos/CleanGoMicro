package tables

import (
	"modules/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	User_id  int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Login    string `json:"login" example:"admin"`
	Password string `json:"password" example:"123456"`
}

func (instance *User) RecordCreate(db models.BaseDatabase) models.Response {
	return db.CreateData(instance)
}
func (instance *User) RecordShow(db models.BaseDatabase) models.Response {
	err := db.ShowData(instance)
	return err
}
func (instance *User) RecordDelete(db models.BaseDatabase) models.Response {
	return db.DeleteData(instance)
}
func (instance *User) RecordUpdate(db models.BaseDatabase) models.Response {
	return db.UpdateData(instance)
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
