package database

import (
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2/log"
)

type TaskDao struct{}

func (currentlDB *TaskDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := currentlDB.convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseUser{}.BadCreate()
	}

	if result := dbConnect.Instance.Create(&user); result.Error != nil {
		log.Debug("create record error!")
		return responses.ResponseUser{}.BadCreate()
	}

	return responses.ResponseUser{}.GoodCreate()
}

func (currentlDB *TaskDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := user.GetId()

	dbConnect := currentlDB.convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseUser{}.BadCreate()
	}

	if result := dbConnect.Instance.First(&user, id); result.Error != nil {
		log.Debug("update record error!")
		return responses.ResponseUser{}.BadUpdate()
	}
	return responses.ResponseUser{}.GoodUpdate()
}

func (currentlDB *TaskDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := user.GetId()

	dbConnect := currentlDB.convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseUser{}.BadCreate()
	}

	result := dbConnect.Instance.Delete(&user, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return responses.ResponseUser{}.BadDelete()
	}
	return responses.ResponseUser{}.GoodDelete()
}

func (currentlDB *TaskDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []tables.User

	dbConnect := currentlDB.convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseUser{}.BadCreate()
	}

	results := dbConnect.Instance.Find(&finded, user)
	if results.Error != nil {
		log.Debug("show record error!")
		return responses.ResponseUser{}.BadShow()
	}

	return responses.ResponseUser{}.GoodShow(finded)
}

func (currentlDB *TaskDao) getData(temp models.Table) (tables.Task, models.Response) {
	task, ok := temp.(*tables.Task)
	if ok == false {
		return tables.Task{}, responses.ResponseUser{}.BadCreate()
	}
	return *task, nil
}

func (currentlDB *TaskDao) convertToPostgres(interf models.DatabaseCore) *PostgresDatabase {
	dbConnect, err := interf.(*PostgresDatabase)
	if !err {
		return nil
	}

	return dbConnect
}
