package database

import (
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2/log"
)

func (currentlDB *PostgresDatabase) CreateData(data models.Table) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	if result := currentlDB.Instance.Create(&user); result.Error != nil {
		log.Debug("create record error!")
		return responses.ResponseUser{}.BadCreate()
	}

	return responses.ResponseUser{}.GoodCreate()
}

func (currentlDB *PostgresDatabase) UpdateData(data models.Table) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := user.GetId()

	if result := currentlDB.Instance.First(&user, id); result.Error != nil {
		log.Debug("update record error!")
		return responses.ResponseUser{}.BadUpdate()
	}
	return responses.ResponseUser{}.GoodUpdate()
}

func (currentlDB *PostgresDatabase) DeleteData(data models.Table) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := user.GetId()
	result := currentlDB.Instance.Delete(&user, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return responses.ResponseUser{}.BadDelete()
	}
	return responses.ResponseUser{}.GoodDelete()
}

func (currentlDB *PostgresDatabase) ShowData(data models.Table) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []tables.User

	results := currentlDB.Instance.Find(&finded, user)
	if results.Error != nil {
		log.Debug("show record error!")
		return responses.ResponseUser{}.BadShow()
	}

	return responses.ResponseUser{}.GoodShow(finded)
}
