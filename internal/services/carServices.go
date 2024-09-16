package services

import (
	"modules/internal/database"
	"modules/internal/database/dao"
	"modules/internal/models"
	"modules/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

// вызов метода внутри соответствующей таблицы, и отправка туда нужного коннекта и дао. Вызывается из роутов.

func CarInsert(instance tables.Car) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordCreate(database.GlobalMongo, dao.GlobalCarDao)
}

func CarShow(instance tables.Car) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordShow(database.GlobalMongo, dao.GlobalCarDao)
}

func CarUpdate(instance tables.Car) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordUpdate(database.GlobalMongo, dao.GlobalCarDao)
}

func CarDelete(instance tables.Car) models.Response {
	log.Debug("services get = ", instance)
	return instance.RecordDelete(database.GlobalMongo, dao.GlobalCarDao)
}
