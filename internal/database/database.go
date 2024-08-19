package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var GlobalComputer ComputerDatabase
var GlobalUser UserDatabase

type ComputerDatabase struct {
	Instance *mongo.Client
}
type UserDatabase struct {
	Instance *gorm.DB
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
