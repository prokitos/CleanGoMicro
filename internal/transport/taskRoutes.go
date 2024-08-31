package transport

import (
	"modules/internal/models/tables"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

func insertTask(c *fiber.Ctx) error {
	var curTask tables.Task
	curTask.Description = "hello world"
	curTask.DueDate = "10.10.2010"
	return services.TaskInsert(curTask).GetError(c)
}
