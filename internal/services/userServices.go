package services

import (
	"modules/internal/database"
	"modules/internal/models"
	"modules/internal/models/tables"
)

func UserInsert(instance tables.User) models.Response {
	return instance.RecordCreate(database.GlobalPostgres)
}

func UserShow(instance tables.User) models.Response {
	return instance.RecordShow(database.GlobalPostgres)
}

func UserUpdate(instance tables.User) models.Response {
	return instance.RecordUpdate(database.GlobalPostgres)
}

func UserDelete(instance tables.User) models.Response {
	return instance.RecordDelete(database.GlobalPostgres)
}
