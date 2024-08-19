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

func (currentlDB *UserDatabase) OpenConnection(config config.PostgresConfig) {
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, config.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		panic("not connection to db")
	}

	currentlDB.Instance = db
}

func (currentlDB *UserDatabase) GlobalSet() {
	GlobalUser = *currentlDB
}

func (currentlDB *UserDatabase) CreateData() models.Response {
	return models.ResponseUser{}.BadCreate()
}

func (currentlDB *UserDatabase) UpdateData() models.Response {
	return models.ResponseUser{}.BadCreate()
}

func (currentlDB *UserDatabase) DeleteData() models.Response {
	return models.ResponseUser{}.BadCreate()
}

func (currentlDB *UserDatabase) ShowData() ([]models.Table, models.Response) {
	return nil, models.ResponseUser{}.BadShow()
}
