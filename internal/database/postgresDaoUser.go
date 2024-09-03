package database

import (
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2/log"
)

type UserDao struct{}

func (currentlDB *UserDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseUser{}.InternalError()
	}

	if result := dbConnect.Instance.Create(&user); result.Error != nil {
		log.Debug("create record error!")
		return responses.ResponseUser{}.BadCreate()
	}

	return responses.ResponseUser{}.GoodCreate()
}

func (currentlDB *UserDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseUser{}.InternalError()
	}

	if result := dbConnect.Instance.Updates(&user); result.Error != nil {
		log.Debug("update record error!")
		return responses.ResponseUser{}.BadUpdate()
	}
	return responses.ResponseUser{}.GoodUpdate()
}

func (currentlDB *UserDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := user.GetId()

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseUser{}.InternalError()
	}

	result := dbConnect.Instance.Delete(&user, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return responses.ResponseUser{}.BadDelete()
	}
	return responses.ResponseUser{}.GoodDelete()
}

func (currentlDB *UserDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []tables.User

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseUser{}.InternalError()
	}

	results := dbConnect.Instance.Find(&finded, user)
	if results.Error != nil {
		log.Debug("show record error!")
		return responses.ResponseUser{}.BadShow()
	}

	return responses.ResponseUser{}.GoodShow(finded)
}

func (currentlDB *UserDao) getData(temp models.Table) (tables.User, models.Response) {
	person, ok := temp.(*tables.User)
	if ok == false {
		return tables.User{}, responses.ResponseUser{}.InternalError()
	}
	return *person, nil
}
