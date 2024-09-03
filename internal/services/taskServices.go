package services

import (
	"modules/internal/database"
	"modules/internal/models"
	"modules/internal/models/tables"
)

func TaskInsert(instance tables.Task) models.Response {
	return instance.RecordCreate(database.GlobalPostgres, database.GlobalTaskDao)
}

func TaskShow(instance tables.Task) models.Response {
	return instance.RecordShow(database.GlobalPostgres, database.GlobalTaskDao)
}

func TaskUpdate(instance tables.Task) models.Response {
	return instance.RecordUpdate(database.GlobalPostgres, database.GlobalTaskDao)
}

func TaskDelete(instance tables.Task) models.Response {
	return instance.RecordDelete(database.GlobalPostgres, database.GlobalTaskDao)
}
