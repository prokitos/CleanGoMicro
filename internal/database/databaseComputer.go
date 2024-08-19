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

func (currentlDB *ComputerDatabase) ConnectToDB(config config.MongoConfig) {

	connectStr := fmt.Sprintf("mongodb://%s:%s", config.Host, config.Port)
	client, err := mongo.NewClient(options.Client().ApplyURI(connectStr))
	checkError(err)

	err = client.Connect(context.TODO())
	checkError(err)

	err = client.Ping(context.TODO(), nil)
	checkError(err)

	currentlDB.Instance = client
}

func (currentlDB *ComputerDatabase) GlobalSet() {
	GlobalComputer = *currentlDB
}

func (currentlDB *ComputerDatabase) CreateData() models.Response {
	return models.ResponseComputer{}.BadShow()
}

func (currentlDB *ComputerDatabase) UpdateData() models.Response {
	return models.ResponseComputer{}.BadShow()
}

func (currentlDB *ComputerDatabase) DeleteData() models.Response {
	return models.ResponseComputer{}.BadShow()
}

func (currentlDB *ComputerDatabase) ShowData() ([]models.Table, models.Response) {
	return nil, models.ResponseComputer{}.BadShow()
}
