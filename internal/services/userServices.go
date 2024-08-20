package services

import (
	"modules/internal/database"
	"modules/internal/models"
)

func UserInsert(instance models.User) models.Response {
	return instance.RecordCreate(database.GlobalUser)
}

func UserShow(instance models.User) models.Response {
	return instance.RecordShow(database.GlobalUser)
}

func UserUpdate(instance models.User) models.Response {
	return instance.RecordUpdate(database.GlobalUser)
}

func UserDelete(instance models.User) models.Response {
	return instance.RecordDelete(database.GlobalUser)
}

// var builder models.TableBuilder
// var temp []models.Table
// temp = append(temp, builder.UserCreate("vanya", "123456").Instance)
// temp = append(temp, builder.UserCreate("seg", "wey").Instance)
