package services

import (
	"modules/internal/database"
	"modules/internal/database/dao"
	"modules/internal/models"
	"modules/internal/models/tables"
)

// вызов метода внутри соответствующей таблицы, и отправка туда нужного коннекта и дао. Вызывается из роутов.

func TaskInsert(instance tables.Task) models.Response {
	return instance.RecordCreate(database.GlobalPostgres, dao.GlobalTaskDao)
}

func TaskShow(instance tables.Task) models.Response {
	return instance.RecordShow(database.GlobalPostgres, dao.GlobalTaskDao)
}

func TaskUpdate(instance tables.Task) models.Response {
	return instance.RecordUpdate(database.GlobalPostgres, dao.GlobalTaskDao)
}

func TaskDelete(instance tables.Task) models.Response {
	return instance.RecordDelete(database.GlobalPostgres, dao.GlobalTaskDao)
}
