package services

import (
	"modules/internal/database"
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"
)

func CarInsert(instance tables.Car) models.Response {
	return instance.RecordCreate(database.GlobalPostgres)
}

func CarShow(instance tables.Car) models.Response {
	return responses.ResponseUser{}.BadCreate()
}

func CarUpdate(instance tables.Car) models.Response {
	return responses.ResponseUser{}.BadCreate()
}

func CarDelete(instance tables.Car) models.Response {
	return responses.ResponseUser{}.BadCreate()
}
