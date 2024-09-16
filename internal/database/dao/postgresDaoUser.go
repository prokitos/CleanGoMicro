package dao

import (
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

// (дао/круд) для таблицы User. вызывается из соответсвующей таблицы.

type UserDao struct{}

// функция которая возвращает респонс текущего дао. Нужен чтобы не менять кучу респонсов у новых дао.
func (currentlDB *UserDao) curResponse() responses.ResponseUser {
	return responses.ResponseUser{}
}

func (currentlDB *UserDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao ", core, " get = ", data)

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Create(&user); result.Error != nil {
		log.Debug("create record error!")
		return currentlDB.curResponse().BadCreate()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodCreate()
}

func (currentlDB *UserDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao ", core, " get = ", data)

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := user.GetId()

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	result := dbConnect.Instance.Delete(&user, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return currentlDB.curResponse().BadDelete()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodDelete()
}

func (currentlDB *UserDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao ", core, " get = ", data)

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Updates(&user); result.Error != nil {
		log.Debug("update record error!")
		return currentlDB.curResponse().BadUpdate()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodUpdate()
}

func (currentlDB *UserDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao ", core, " get = ", data)

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []tables.User

	dbConnect := convertToPostgres(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	results := dbConnect.Instance.Find(&finded, user)
	if results.Error != nil || results.RowsAffected == 0 {
		log.Debug("show record error!")
		return currentlDB.curResponse().BadShow()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodShow(finded)
}

// перево интерфейса таблицы в конкретную таблицу
func (currentlDB *UserDao) getData(temp models.Table) (tables.User, models.Response) {
	person, ok := temp.(*tables.User)
	if ok == false {
		return tables.User{}, currentlDB.curResponse().InternalError()
	}
	return *person, nil
}
