package models

import "github.com/gofiber/fiber/v2"

type Response interface {
	GetError(c *fiber.Ctx) error
	GoodCreate() Response
	BadCreate() Response
	GoodShow([]Table) Response
	BadShow() Response
	Validate() bool
}

type Table interface {
	RecordCreate()
	RecordDelete()
	RecordShow()
	RecordUpdate()
}
type GormDatabase interface {
	OpenConnection()
	StartMigration()
	GlobalSet()
	CreateData()
	DeleteData()
	UpdateData()
	ShowData()
}
