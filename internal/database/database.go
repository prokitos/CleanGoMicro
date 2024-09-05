package database

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// глобальные переменные для хранения всех подключений к бд.

var GlobalMongo *MongoDatabase
var GlobalPostgres *PostgresDatabase

type MongoDatabase struct {
	Instance *mongo.Client
}
type PostgresDatabase struct {
	Instance *gorm.DB
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
