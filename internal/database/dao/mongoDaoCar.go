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

func (currentlDB *CarDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {

	curCar, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return responses.ResponseCar{}.InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("cars")
	_, err := collection.InsertOne(context.TODO(), curCar)
	if err != nil {
		return responses.ResponseCar{}.BadCreate()
	}

	return responses.ResponseCar{}.GoodCreate()
}

func (currentlDB *CarDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {

	curCar, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := curCar.OutputGet()

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return responses.ResponseCar{}.InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("cars")
	deleteResult, _ := collection.DeleteOne(context.TODO(), temp)
	if deleteResult.DeletedCount == 0 {
		return responses.ResponseCar{}.BadDelete()
	}

	return responses.ResponseCar{}.GoodDelete()
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
		return responses.ResponseCar{}.InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("cars")
	_, err := collection.UpdateMany(context.TODO(), bson.D{{"_id", temp.Car_id}}, update)
	if err != nil {
		return responses.ResponseCar{}.BadUpdate()
	}

	return responses.ResponseCar{}.GoodUpdate()
}

func (currentlDB *CarDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {

	curCar, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := curCar.OutputGet()

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return responses.ResponseCar{}.InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("cars")
	cur, err := collection.Find(context.TODO(), temp)
	if err != nil {
		return responses.ResponseCar{}.BadShow()
	}
	if cur.RemainingBatchLength() == 0 {
		return responses.ResponseCar{}.BadShow()
	}

	var finded []tables.Car
	for cur.Next(context.TODO()) {
		var elem tables.Car
		err := cur.Decode(&elem)
		if err != nil {
			return responses.ResponseCar{}.InternalError()
		}
		finded = append(finded, elem)
	}

	return responses.ResponseCar{}.GoodShow(finded)
}

func (currentlDB *CarDao) getData(temp models.Table) (tables.Car, models.Response) {
	cars, ok := temp.(*tables.Car)
	if ok == false {
		return tables.Car{}, responses.ResponseCar{}.InternalError()
	}
	return *cars, nil
}
