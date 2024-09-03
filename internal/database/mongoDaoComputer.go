package database

import (
	"context"
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"go.mongodb.org/mongo-driver/bson"
)

type ComputerDao struct{}

func (currentlDB *ComputerDao) CreateData(data models.Table, core models.DatabaseCore) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return responses.ResponseComputer{}.InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("comps")
	_, err := collection.InsertOne(context.TODO(), computer)
	checkError(err)

	return responses.ResponseComputer{}.GoodCreate()
}

func (currentlDB *ComputerDao) DeleteData(data models.Table, core models.DatabaseCore) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return responses.ResponseComputer{}.InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("comps")
	deleteResult, _ := collection.DeleteOne(context.TODO(), temp)
	if deleteResult.DeletedCount == 0 {
		return responses.ResponseComputer{}.BadDelete()
	}

	return responses.ResponseComputer{}.GoodDelete()
}

func (currentlDB *ComputerDao) UpdateData(data models.Table, core models.DatabaseCore) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()
	update := bson.D{{"$set", temp}}

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return responses.ResponseComputer{}.InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("comps")
	_, err := collection.UpdateMany(context.TODO(), bson.D{{"_id", temp.Computer_id}}, update)
	checkError(err)

	return responses.ResponseComputer{}.GoodUpdate()
}

func (currentlDB *ComputerDao) ShowData(data models.Table, core models.DatabaseCore) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()

	dbConnect := convertToMongo(core)
	if dbConnect == nil {
		return responses.ResponseComputer{}.InternalError()
	}

	collection := dbConnect.Instance.Database("test").Collection("comps")
	cur, err := collection.Find(context.TODO(), temp)
	checkError(err)
	var finded []tables.Computer

	for cur.Next(context.TODO()) {
		var elem tables.Computer
		err := cur.Decode(&elem)
		checkError(err)
		finded = append(finded, elem)
	}

	return responses.ResponseComputer{}.GoodShow(finded)
}

func (currentlDB *ComputerDao) getData(temp models.Table) (tables.Computer, models.Response) {
	devices, ok := temp.(*tables.Computer)
	if ok == false {
		return tables.Computer{}, responses.ResponseComputer{}.InternalError()
	}
	return *devices, nil
}
