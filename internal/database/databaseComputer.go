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

func (currentlDB *ComputerDatabase) CreateData(computer models.Table) models.Response {
	return models.ResponseComputer{}.BadShow()
}

func (currentlDB *ComputerDatabase) UpdateData(computer models.Table) models.Response {
	return models.ResponseComputer{}.BadShow()
}

func (currentlDB *ComputerDatabase) DeleteData(computer models.Table) models.Response {
	return models.ResponseComputer{}.BadShow()
}

func (currentlDB *ComputerDatabase) ShowData(computer models.Table) models.Response {
	return models.ResponseComputer{}.BadShow()
}
