package dao

import (
	"modules/internal/database"
	"modules/internal/models"
)

// Глобальные переменные которые хранят все существующие дао, а также конвертация интерфейса подключения к базе данных в конкретное подключение.

var GlobalUserDao *UserDao
var GlobalTaskDao *TaskDao
var GlobalComputerDao *ComputerDao
var GlobalCarDao *CarDao

func convertToMongo(interf models.DatabaseCore) *database.MongoDatabase {
	dbConnect, err := interf.(*database.MongoDatabase)
	if !err {
		return nil
	}

	return dbConnect
}

func convertToPostgres(interf models.DatabaseCore) *database.PostgresDatabase {
	dbConnect, err := interf.(*database.PostgresDatabase)
	if !err {
		return nil
	}

	return dbConnect
}
