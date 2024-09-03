package transport

import (
	"modules/internal/models/tables"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

func getTask(c *fiber.Ctx) error {
	var curTask tables.Task
	curTask.GetQueryParams(c)
	curTask.GetQueryId(c)
	return services.TaskShow(curTask).GetError(c)
}

func insertTask(c *fiber.Ctx) error {
	var curTask tables.Task
	curTask.GetQueryParams(c)
	return services.TaskInsert(curTask).GetError(c)
}

func deleteTask(c *fiber.Ctx) error {
	var curTask tables.Task
	curTask.GetQueryId(c)
	return services.TaskDelete(curTask).GetError(c)
}

func updateTask(c *fiber.Ctx) error {
	var curTask tables.Task
	curTask.GetQueryParams(c)
	curTask.GetQueryId(c)
	return services.TaskUpdate(curTask).GetError(c)
}
