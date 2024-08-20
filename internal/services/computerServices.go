package services

import (
	"modules/internal/database"
	"modules/internal/models"
)

func ComputerInsert(instance models.Computer) models.Response {
	return instance.RecordCreate(database.GlobalComputer)
}

func ComputerShow(instance models.Computer) models.Response {
	return instance.RecordShow(database.GlobalComputer)
}

func ComputerUpdate(instance models.Computer) models.Response {
	return instance.RecordUpdate(database.GlobalComputer)
}

func ComputerDelete(instance models.Computer) models.Response {
	return instance.RecordDelete(database.GlobalComputer)
}
