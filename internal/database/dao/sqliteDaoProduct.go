package dao

import (
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	log "github.com/sirupsen/logrus"
)

// (дао/круд) для таблицы Product. вызывается из соответсвующей таблицы.

type ProductDao struct{}

// функция которая возвращает респонс текущего дао. Нужен чтобы не менять кучу респонсов у новых дао.
func (currentlDB *ProductDao) curResponse() responses.ResponseProduct {
	return responses.ResponseProduct{}
}

func (currentlDB *ProductDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {

	prod, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToSqlite(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Create(&prod); result.Error != nil {
		log.Debug("create record error!")
		return currentlDB.curResponse().BadCreate()
	}

	return currentlDB.curResponse().GoodCreate()
}

func (currentlDB *ProductDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {

	prod, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := prod.GetId()

	dbConnect := convertToSqlite(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	result := dbConnect.Instance.Delete(&prod, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return currentlDB.curResponse().BadDelete()
	}
	return currentlDB.curResponse().GoodDelete()
}

func (currentlDB *ProductDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {

	prod, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToSqlite(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	if result := dbConnect.Instance.Updates(&prod); result.Error != nil {
		log.Debug("update record error!")
		return currentlDB.curResponse().BadUpdate()
	}
	return currentlDB.curResponse().GoodUpdate()
}

func (currentlDB *ProductDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {

	prod, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []tables.Product

	dbConnect := convertToSqlite(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	results := dbConnect.Instance.Find(&finded, prod)
	if results.Error != nil || results.RowsAffected == 0 {
		log.Debug("show record error!")
		return currentlDB.curResponse().BadShow()
	}

	return currentlDB.curResponse().GoodShow(finded)
}

// перево интерфейса таблицы в конкретную таблицу
func (currentlDB *ProductDao) getData(temp models.Table) (tables.Product, models.Response) {
	prod, ok := temp.(*tables.Product)
	if ok == false {
		return tables.Product{}, currentlDB.curResponse().InternalError()
	}
	return *prod, nil
}
