package dao

import (
	"context"
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"go.mongodb.org/mongo-driver/bson"
)

// (дао/круд) для таблицы Car. вызывается из соответсвующей таблицы.

type CarDao struct{}

// функция которая возвращает респонс текущего дао. Нужен чтобы не менять кучу респонсов у новых дао.
func (currentlDB *CarDao) curResponse() responses.ResponseCar {
	return responses.ResponseCar{}
}

func (currentlDB *CarDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {

	curCar, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("cars")
	_, err := collection.InsertOne(context.TODO(), curCar)
	if err != nil {
		return currentlDB.curResponse().BadCreate()
	}

	return currentlDB.curResponse().GoodCreate()
}

func (currentlDB *CarDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {

	curCar, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := curCar.OutputGet()

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("cars")
	deleteResult, _ := collection.DeleteOne(context.TODO(), temp)
	if deleteResult.DeletedCount == 0 {
		return currentlDB.curResponse().BadDelete()
	}

	return currentlDB.curResponse().GoodDelete()
}

func (currentlDB *CarDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {

	curCar, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := curCar.OutputGet()
	update := bson.D{{"$set", temp}}

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("cars")
	_, err := collection.UpdateMany(context.TODO(), bson.D{{"_id", temp.Car_id}}, update)
	if err != nil {
		return currentlDB.curResponse().BadUpdate()
	}

	return currentlDB.curResponse().GoodUpdate()
}

func (currentlDB *CarDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {

	curCar, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := curCar.OutputGet()

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return currentlDB.curResponse().InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("cars")
	cur, err := collection.Find(context.TODO(), temp)
	if err != nil {
		return currentlDB.curResponse().BadShow()
	}
	if cur.RemainingBatchLength() == 0 {
		return currentlDB.curResponse().BadShow()
	}

	var finded []tables.Car
	for cur.Next(context.TODO()) {
		var elem tables.Car
		err := cur.Decode(&elem)
		if err != nil {
			return currentlDB.curResponse().InternalError()
		}
		finded = append(finded, elem)
	}

	return currentlDB.curResponse().GoodShow(finded)
}

// перево интерфейса таблицы в конкретную таблицу
func (currentlDB *CarDao) getData(temp models.Table) (tables.Car, models.Response) {
	cars, ok := temp.(*tables.Car)
	if ok == false {
		return tables.Car{}, currentlDB.curResponse().InternalError()
	}
	return *cars, nil
}
