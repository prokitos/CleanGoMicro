package database

import (
	"context"
	"fmt"
	"modules/internal/config"
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (currentlDB *MongoDatabase) Run(config config.MainConfig) {
	currentlDB.OpenConnection(config)
	currentlDB.StartMigration()
	currentlDB.GlobalSet()
}

func (currentlDB *MongoDatabase) StartMigration() {
	log.Debug("migration dont exist")
}

func (currentlDB *MongoDatabase) OpenConnection(config config.MainConfig) {

	connectStr := fmt.Sprintf("mongodb://%s:%s", config.MongoDB.Host, config.MongoDB.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectStr))
	checkError(err)

	err = client.Connect(context.TODO())
	checkError(err)

	err = client.Ping(context.TODO(), nil)
	checkError(err)

	currentlDB.Instance = client
}

func (currentlDB *MongoDatabase) GlobalSet() {
	GlobalMongo = currentlDB
}

func (currentlDB *MongoDatabase) getData(temp models.Table) (tables.Computer, models.Response) {
	devices, ok := temp.(*tables.Computer)
	if ok == false {
		return tables.Computer{}, responses.ResponseUser{}.BadCreate()
	}
	return *devices, nil
}
