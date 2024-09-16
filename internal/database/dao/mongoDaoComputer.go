package dao

import (
	"context"
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
)

// (дао/круд) для таблицы Computer. вызывается из соответсвующей таблицы.

type ComputerDao struct{}

// функция которая возвращает респонс текущего дао. Нужен чтобы не менять кучу респонсов у новых дао.
func (currentlDB *ComputerDao) curResponse() responses.ResponseComputer {
	return responses.ResponseComputer{}
}

func (currentlDB *ComputerDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao ", core, " get = ", data)

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("comps")
	_, err := collection.InsertOne(context.TODO(), computer)
	if err != nil {
		return currentlDB.curResponse().BadCreate()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodCreate()
}

func (currentlDB *ComputerDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao ", core, " get = ", data)

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("comps")
	deleteResult, _ := collection.DeleteOne(context.TODO(), temp)
	if deleteResult.DeletedCount == 0 {
		return currentlDB.curResponse().BadDelete()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodDelete()
}

func (currentlDB *ComputerDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao ", core, " get = ", data)

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()
	update := bson.D{{"$set", temp}}

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("comps")
	_, err := collection.UpdateMany(context.TODO(), bson.D{{"_id", temp.Computer_id}}, update)
	if err != nil {
		return currentlDB.curResponse().BadUpdate()
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodUpdate()
}

func (currentlDB *ComputerDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {
	log.Debug("dao ", core, " get = ", data)

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("comps")
	cur, err := collection.Find(context.TODO(), temp)
	if err != nil {
		return currentlDB.curResponse().BadShow()
	}
	if cur.RemainingBatchLength() == 0 {
		return responses.ResponseCar{}.BadShow()
	}

	var finded []tables.Computer
	for cur.Next(context.TODO()) {
		var elem tables.Computer
		err := cur.Decode(&elem)
		if err != nil {
			return currentlDB.curResponse().InternalError()
		}
		finded = append(finded, elem)
	}

	log.Debug("dao complete")
	return currentlDB.curResponse().GoodShow(finded)
}

// перево интерфейса таблицы в конкретную таблицу
func (currentlDB *ComputerDao) getData(temp models.Table) (tables.Computer, models.Response) {
	devices, ok := temp.(*tables.Computer)
	if ok == false {
		return tables.Computer{}, currentlDB.curResponse().InternalError()
	}
	return *devices, nil
}
