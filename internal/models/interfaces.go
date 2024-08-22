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
	RecordCreate(GormDatabase) Response
	RecordDelete(GormDatabase) Response
	RecordShow(GormDatabase) Response
	RecordUpdate(GormDatabase) Response
	GetId() int
}
type GormDatabase interface {
	OpenConnection(config.MainConfig)
	StartMigration()
	GlobalSet()
	CreateData(Table) Response
	DeleteData(Table) Response
	UpdateData(Table) Response
	ShowData(Table) Response
}
