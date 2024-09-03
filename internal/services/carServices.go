package services

import (
	"modules/internal/database"
	"modules/internal/models"
	"modules/internal/models/tables"
)

func CarInsert(instance tables.Car) models.Response {
	return instance.RecordCreate(database.GlobalMongo, database.GlobalCarDao)
}

func CarShow(instance tables.Car) models.Response {
	return instance.RecordShow(database.GlobalMongo, database.GlobalCarDao)
}

func CarUpdate(instance tables.Car) models.Response {
	return instance.RecordUpdate(database.GlobalMongo, database.GlobalCarDao)
}

func CarDelete(instance tables.Car) models.Response {
	return instance.RecordDelete(database.GlobalMongo, database.GlobalCarDao)
}
