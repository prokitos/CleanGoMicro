package services

import (
	"modules/internal/models"
	"modules/internal/models/responses"
	"modules/internal/models/tables"
)

func TaskInsert(instance tables.Task) models.Response {
	return responses.ResponseUser{}.BadCreate()
}

func TaskShow(instance tables.Task) models.Response {
	return responses.ResponseUser{}.BadCreate()
}

func TaskUpdate(instance tables.Task) models.Response {
	return responses.ResponseUser{}.BadCreate()
}

func TaskDelete(instance tables.Task) models.Response {
	return responses.ResponseUser{}.BadCreate()
}
