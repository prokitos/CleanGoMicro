package database

import (
	"context"
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"go.mongodb.org/mongo-driver/bson"
)

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
	checkError(err)

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
	checkError(err)

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
	checkError(err)
	var finded []tables.Car

	for cur.Next(context.TODO()) {
		var elem tables.Car
		err := cur.Decode(&elem)
		checkError(err)
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
