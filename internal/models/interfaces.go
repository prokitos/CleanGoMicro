package models

import (
	"modules/internal/config"

	"github.com/gofiber/fiber/v2"
)

type Response interface {
	GetError(c *fiber.Ctx) error
	GoodCreate() Response
	BadCreate() Response
	GoodShow([]Table) Response
	BadShow() Response
	Validate() bool
}

type Table interface {
	RecordCreate(GormDatabase) Response
	RecordDelete(GormDatabase) Response
	RecordShow(GormDatabase) Response
	RecordUpdate(GormDatabase) Response
}
type GormDatabase interface {
	OpenConnection(config.MainConfig)
	StartMigration()
	GlobalSet()
	CreateData(Table) Response
	DeleteData(Table) Response
	UpdateData(Table) Response
	ShowData(Table) ([]Table, Response)
}
