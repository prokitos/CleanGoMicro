package database

import (
	"context"
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"go.mongodb.org/mongo-driver/bson"
)

func (currentlDB *MongoDatabase) CreateData(data models.Table) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	collection := currentlDB.Instance.Database("test").Collection("testovs")
	_, err := collection.InsertOne(context.TODO(), computer)
	checkError(err)

	return responses.ResponseComputer{}.GoodCreate()
}

func (currentlDB *MongoDatabase) DeleteData(data models.Table) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()

	collection := currentlDB.Instance.Database("test").Collection("testovs")
	deleteResult, _ := collection.DeleteOne(context.TODO(), temp)
	if deleteResult.DeletedCount == 0 {
		return responses.ResponseComputer{}.BadDelete()
	}

	return responses.ResponseComputer{}.GoodDelete()
}

func (currentlDB *MongoDatabase) UpdateData(data models.Table) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()
	update := bson.D{{"$set", temp}}

	collection := currentlDB.Instance.Database("test").Collection("testovs")
	_, err := collection.UpdateMany(context.TODO(), bson.D{{"_id", temp.Computer_id}}, update)
	checkError(err)

	return responses.ResponseComputer{}.GoodUpdate()
}

func (currentlDB *MongoDatabase) ShowData(data models.Table) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	temp := computer.OutputGet()

	collection := currentlDB.Instance.Database("test").Collection("testovs")
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

func (currentlDB *MongoDatabase) getData(temp models.Table) (tables.Computer, models.Response) {
	devices, ok := temp.(*tables.Computer)
	if ok == false {
		return tables.Computer{}, responses.ResponseComputer{}.BadCreate()
	}
	return *devices, nil
}
