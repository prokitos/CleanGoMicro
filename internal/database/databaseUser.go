package database

import (
	"fmt"
	"modules/internal/config"
	"modules/internal/models"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (currentlDB *UserDatabase) StartMigration() {
	currentlDB.Instance.AutoMigrate(models.User{})
	log.Debug("migration complete")
}

func (currentlDB *UserDatabase) OpenConnection(config config.MainConfig) {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.PostgresDB.User, config.PostgresDB.Pass, config.PostgresDB.Host, config.PostgresDB.Port, config.PostgresDB.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		panic("not connection to db")
	}

	currentlDB.Instance = db
}

func (currentlDB *UserDatabase) GlobalSet() {
	GlobalUser = *currentlDB
}

func (currentlDB *UserDatabase) CreateData(user models.User) models.Response {
	return models.ResponseUser{}.BadCreate()
}

func (currentlDB *UserDatabase) UpdateData(user models.User) models.Response {
	return models.ResponseUser{}.BadCreate()
}

func (currentlDB *UserDatabase) DeleteData(user models.User) models.Response {
	return models.ResponseUser{}.BadCreate()
}

func (currentlDB *UserDatabase) ShowData(user models.User) ([]models.Table, models.Response) {
	return nil, models.ResponseUser{}.BadShow()
}
