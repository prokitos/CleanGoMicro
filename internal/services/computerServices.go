package services

import (
	"modules/internal/database"
	"modules/internal/models"
	"modules/internal/models/tables"
)

func ComputerInsert(instance tables.Computer) models.Response {
	return instance.RecordCreate(database.GlobalMongo, database.GlobalComputerDao)
}

func ComputerShow(instance tables.Computer) models.Response {
	return instance.RecordShow(database.GlobalMongo, database.GlobalComputerDao)
}

func ComputerUpdate(instance tables.Computer) models.Response {
	return instance.RecordUpdate(database.GlobalMongo, database.GlobalComputerDao)
}

func ComputerDelete(instance tables.Computer) models.Response {
	return instance.RecordDelete(database.GlobalMongo, database.GlobalComputerDao)
}
