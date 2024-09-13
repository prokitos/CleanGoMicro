package database

import (
	"modules/internal/config"
	"modules/internal/models/tables"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// запуск, соединение и миграция для postgresDB. (если будут подключения к нескольким базам postgres, то создавать ещё файлы, и делать им названия postgresNameDatabase)

func (currentlDB *SqliteDatabase) Run(config config.MainConfig) {
	currentlDB.OpenConnection(config)
	currentlDB.StartMigration()
	currentlDB.GlobalSet()
}

func (currentlDB *SqliteDatabase) StartMigration() {
	currentlDB.Instance.AutoMigrate(tables.Product{})
	log.Debug("migration complete")
}

func (currentlDB *SqliteDatabase) OpenConnection(config config.MainConfig) {

	db, err := gorm.Open(sqlite.Open(config.SqliteDb.Name), &gorm.Config{})
	if err != nil {
		panic("not connection to db")
	}

	currentlDB.Instance = db

}

func (currentlDB *SqliteDatabase) GlobalSet() {
	GlobalSqlite = currentlDB
}
