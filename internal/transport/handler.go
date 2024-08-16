package transport

import (
	"github.com/gofiber/fiber/v2"
)

func SetHandlers(instance *fiber.App) {
	instance.Get("/user", getUser)
	instance.Post("/user", insertUser)

	instance.Get("/computer", getComputer)
	instance.Post("/computer", insertComputer)
}
