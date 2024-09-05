package services

import (
	"modules/internal/database"
	"modules/internal/database/dao"
	"modules/internal/models"
	"modules/internal/models/tables"
)

// вызов метода внутри соответствующей таблицы, и отправка туда нужного коннекта и дао. Вызывается из роутов.

func ComputerInsert(instance tables.Computer) models.Response {
	return instance.RecordCreate(database.GlobalMongo, dao.GlobalComputerDao)
}

func ComputerShow(instance tables.Computer) models.Response {
	return instance.RecordShow(database.GlobalMongo, dao.GlobalComputerDao)
}

func ComputerUpdate(instance tables.Computer) models.Response {
	return instance.RecordUpdate(database.GlobalMongo, dao.GlobalComputerDao)
}

func ComputerDelete(instance tables.Computer) models.Response {
	return instance.RecordDelete(database.GlobalMongo, dao.GlobalComputerDao)
}
