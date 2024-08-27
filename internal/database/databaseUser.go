package database

import (
	"fmt"
	"modules/internal/config"
	"modules/internal/models"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (currentlDB *UserDatabase) Run(config config.MainConfig) {
	currentlDB.OpenConnection(config)
	currentlDB.StartMigration()
	currentlDB.GlobalSet()
}

func (currentlDB *UserDatabase) StartMigration() {
	currentlDB.Instance.AutoMigrate(models.User{})
	log.Debug("migration complete")
}

func (currentlDB *UserDatabase) OpenConnection(config config.MainConfig) {

	err := currentlDB.checkDatabaseCreated(config)
	if err != nil {
		panic("not connection to db")
	}

	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.PostgresDB.User, config.PostgresDB.Pass, config.PostgresDB.Host, config.PostgresDB.Port, config.PostgresDB.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		panic("not connection to db")
	}

	currentlDB.Instance = db
}

// проверка если есть база данных. если нет, то создать.
func (currentlDB *UserDatabase) checkDatabaseCreated(config config.MainConfig) error {

	// открытие соеднение с базой по стандарту
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.PostgresDB.User, config.PostgresDB.Pass, config.PostgresDB.Host, config.PostgresDB.Port, "postgres")
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Error("database don't open")
		return models.ResponseGlobal{}.InternalError()
	}

	// закрытие бд
	sql, _ := db.DB()
	defer func() {
		_ = sql.Close()
	}()

	// проверка если есть нужная нам база данных
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", config.PostgresDB.Name)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		log.Error("error, dont read bd")
		return models.ResponseGlobal{}.InternalError()
	}

	// если нет, то создать
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", config.PostgresDB.Name)
		if rs := db.Exec(stmt); rs.Error != nil {
			log.Error("error, dont create a database")
			return models.ResponseGlobal{}.InternalError()
		}
	}

	return nil
}

func (currentlDB *UserDatabase) GlobalSet() {
	GlobalUser = currentlDB
}

func (currentlDB *UserDatabase) getData(temp models.Table) (models.User, models.Response) {
	person, ok := temp.(*models.User)
	if ok == false {
		return models.User{}, models.ResponseUser{}.BadCreate()
	}
	return *person, nil
}

func (currentlDB *UserDatabase) CreateData(data models.Table) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	if result := currentlDB.Instance.Create(&user); result.Error != nil {
		log.Debug("create record error!")
		return models.ResponseUser{}.BadCreate()
	}

	return models.ResponseUser{}.GoodCreate()
}

func (currentlDB *UserDatabase) UpdateData(data models.Table) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := user.GetId()

	if result := currentlDB.Instance.First(&user, id); result.Error != nil {
		log.Debug("update record error!")
		return models.ResponseUser{}.BadUpdate()
	}
	return models.ResponseUser{}.GoodUpdate()
}

func (currentlDB *UserDatabase) DeleteData(data models.Table) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	id := user.GetId()
	result := currentlDB.Instance.Delete(&user, id)
	if result.RowsAffected == 0 || result.Error != nil {
		return models.ResponseUser{}.BadDelete()
	}
	return models.ResponseUser{}.GoodDelete()
}

func (currentlDB *UserDatabase) ShowData(data models.Table) models.Response {

	user, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []models.User

	results := currentlDB.Instance.Find(&finded, user)
	if results.Error != nil {
		log.Debug("show record error!")
		return models.ResponseUser{}.BadShow()
	}

	return models.ResponseUser{}.GoodShow(finded)
}
