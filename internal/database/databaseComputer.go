package database

import (
	"context"
	"fmt"
	"modules/internal/config"
	"modules/internal/models"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (currentlDB *ComputerDatabase) StartMigration() {
	log.Debug("migration dont exist")
}

func (currentlDB *ComputerDatabase) OpenConnection(config config.MainConfig) {

	connectStr := fmt.Sprintf("mongodb://%s:%s", config.MongoDB.Host, config.MongoDB.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectStr))
	checkError(err)

	err = client.Connect(context.TODO())
	checkError(err)

	err = client.Ping(context.TODO(), nil)
	checkError(err)

	currentlDB.Instance = client
}

func (currentlDB *ComputerDatabase) GlobalSet() {
	GlobalComputer = currentlDB
}

func (currentlDB *ComputerDatabase) getData(temp models.Table) (models.Computer, models.Response) {
	devices, ok := temp.(*models.Computer)
	if ok == false {
		return models.Computer{}, models.ResponseUser{}.BadCreate()
	}
	return *devices, nil
}

func (currentlDB *ComputerDatabase) CreateData(data models.Table) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	collection := currentlDB.Instance.Database("test").Collection("testovs")
	_, err := collection.InsertOne(context.TODO(), computer)
	checkError(err)

	return models.ResponseComputer{}.GoodCreate()
}

func (currentlDB *ComputerDatabase) UpdateData(data models.Table) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var computer_id models.Computer
	computer_id.Computer_id = computer.Computer_id

	collection := currentlDB.Instance.Database("test").Collection("testovs")
	_, err := collection.UpdateMany(context.TODO(), computer_id, computer)
	checkError(err)

	return models.ResponseComputer{}.GoodUpdate()
}

func (currentlDB *ComputerDatabase) DeleteData(data models.Table) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	collection := currentlDB.Instance.Database("test").Collection("testovs")
	_, err := collection.DeleteMany(context.TODO(), computer)
	checkError(err)

	return models.ResponseComputer{}.GoodDelete()
}

func (currentlDB *ComputerDatabase) ShowData(data models.Table) models.Response {

	computer, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	collection := currentlDB.Instance.Database("test").Collection("testovs")
	cur, err := collection.Find(context.TODO(), computer)
	checkError(err)
	var finded []models.Computer

	for cur.Next(context.TODO()) {
		var elem models.Computer
		err := cur.Decode(&elem)
		checkError(err)
		finded = append(finded, elem)
	}

	return models.ResponseComputer{}.GoodShow(finded)
}
