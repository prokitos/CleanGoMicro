package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type GormDatabase interface {
	OpenConnection()
	StartMigration()
	GlobalSet()
}

var GlobalUser UserDatabase
var GlobalComputer ComputerDatabase

type UserDatabase struct {
	Instance *gorm.DB
}

type ComputerDatabase struct {
	Instance *mongo.Client
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
