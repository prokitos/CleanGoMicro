package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var GlobalMongo *MongoDatabase
var GlobalPostgres *PostgresDatabase

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
