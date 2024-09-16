package database

import (
	"context"
	"fmt"
	"modules/internal/config"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// запуск, соединение и миграция для mongoDB. (если будут подключения к нескольким базам Mongo, то создавать ещё файлы, и делать им названия mongoNameDatabase)

func (currentlDB *MongoDatabase) Run(config config.MainConfig) {
	currentlDB.OpenConnection(config)
	currentlDB.StartMigration()
	currentlDB.GlobalSet()
}

func (currentlDB *MongoDatabase) StartMigration() {
	log.Debug("migration dont exist for this database")
}

func (currentlDB *MongoDatabase) OpenConnection(config config.MainConfig) {

	connectStr := fmt.Sprintf("mongodb://%s:%s", config.MongoDB.Host, config.MongoDB.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectStr))
	CheckError(err)

	err = client.Connect(context.TODO())
	CheckError(err)

	err = client.Ping(context.TODO(), nil)
	CheckError(err)

	currentlDB.Instance = client
}

func (currentlDB *MongoDatabase) GlobalSet() {
	GlobalMongo = currentlDB
}
