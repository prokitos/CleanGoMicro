package models

import (
	"modules/internal/config"

	"github.com/gofiber/fiber/v2"
)

type Response interface {
	GetError(c *fiber.Ctx) error
	Validate() bool
}

type Table interface {
	RecordCreate(BaseDatabase) Response
	RecordDelete(BaseDatabase) Response
	RecordShow(BaseDatabase) Response
	RecordUpdate(BaseDatabase) Response
}
type BaseDatabase interface {
	OpenConnection(config.MainConfig)
	StartMigration()
	GlobalSet()
	CreateData(Table) Response
	DeleteData(Table) Response
	UpdateData(Table) Response
	ShowData(Table) Response
}
