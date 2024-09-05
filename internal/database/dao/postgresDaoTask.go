package dao

import (
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2/log"
)

// (дао/круд) для таблицы Task. вызывается из соответсвующей таблицы.

type TaskDao struct{}

func (currentlDB *TaskDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseTask{}.InternalError()
	}

	if result := dbConnect.Instance.Create(&task); result.Error != nil {
		log.Debug("create record error!")
		return responses.ResponseTask{}.BadCreate()
	}

	return responses.ResponseTask{}.GoodCreate()
}

func (currentlDB *TaskDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := task.GetId()

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseTask{}.InternalError()
	}

	result := dbConnect.Instance.Delete(&task, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return responses.ResponseTask{}.BadDelete()
	}
	return responses.ResponseTask{}.GoodDelete()
}

func (currentlDB *TaskDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseTask{}.InternalError()
	}

	if result := dbConnect.Instance.Updates(&task); result.Error != nil {
		log.Debug("update record error!")
		return responses.ResponseTask{}.BadUpdate()
	}
	return responses.ResponseTask{}.GoodUpdate()
}

func (currentlDB *TaskDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []tables.Task

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return responses.ResponseTask{}.InternalError()
	}

	results := dbConnect.Instance.Find(&finded, task)
	if results.Error != nil || results.RowsAffected == 0 {
		log.Debug("show record error!")
		return responses.ResponseTask{}.BadShow()
	}

	return responses.ResponseTask{}.GoodShow(finded)
}

func (currentlDB *TaskDao) getData(temp models.Table) (tables.Task, models.Response) {
	task, ok := temp.(*tables.Task)
	if ok == false {
		return tables.Task{}, responses.ResponseTask{}.InternalError()
	}
	return *task, nil
}
