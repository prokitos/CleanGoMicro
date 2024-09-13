package services

import (
	"modules/internal/database"
	"modules/internal/database/dao"
	"modules/internal/models"
	"modules/internal/models/tables"
)

// вызов метода внутри соответствующей таблицы, и отправка туда нужного коннекта и дао. Вызывается из роутов.

func ProductInsert(instance tables.Product) models.Response {
	return instance.RecordCreate(database.GlobalSqlite, dao.GlobalProductDao)
}

func ProductShow(instance tables.Product) models.Response {
	return instance.RecordShow(database.GlobalSqlite, dao.GlobalProductDao)
}

func ProductUpdate(instance tables.Product) models.Response {
	return instance.RecordUpdate(database.GlobalSqlite, dao.GlobalProductDao)
}

func ProductDelete(instance tables.Product) models.Response {
	return instance.RecordDelete(database.GlobalSqlite, dao.GlobalProductDao)
}
