package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var GlobalMongo *MongoDatabase
var GlobalPostgres *PostgresDatabase

var GlobalUserDao *UserDao
var GlobalTaskDao *TaskDao
var GlobalComputerDao *ComputerDao
var GlobalCarDao *CarDao

type MongoDatabase struct {
	Instance *mongo.Client
}
type PostgresDatabase struct {
	Instance *gorm.DB
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
