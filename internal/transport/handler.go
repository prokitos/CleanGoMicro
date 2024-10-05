package transport

import (
	"modules/internal/metrics"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// здесь хранятся хэндлеры.

func SetHandlers(instance *fiber.App) {

	reg := prometheus.NewRegistry()
	reg.MustRegister(metrics.RequestDuration)
	reg.MustRegister(metrics.RequestStatus)
	instance.Use(metrics.Observer)

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	go http.ListenAndServe(":8082", mux)

	instance.Get("/user", getUser)
	instance.Post("/user", insertUser)
	instance.Delete("/user", deleteUser)
	instance.Put("/user", updateUser)

	instance.Get("/computer", getComputer)
	instance.Post("/computer", insertComputer)
	instance.Delete("/computer", deleteComputer)
	instance.Put("/computer", updateComputer)

	instance.Get("/car", getCar)
	instance.Post("/car", insertCar)
	instance.Delete("/car", deleteCar)
	instance.Put("/car", updateCar)

	instance.Post("/task", insertTask)
	instance.Get("/task", getTask)
	instance.Delete("/task", deleteTask)
	instance.Put("/task", updateTask)

	instance.Post("/product", insertProduct)
	instance.Get("/product", getProduct)
	instance.Delete("/product", deleteProduct)
	instance.Put("/product", updateProduct)

}
