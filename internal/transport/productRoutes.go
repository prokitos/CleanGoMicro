package transport

import (
	"modules/internal/models/tables"
	"modules/internal/services"

	"github.com/gofiber/fiber/v2"
)

// роуты для product

func getProduct(c *fiber.Ctx) error {
	var curProd tables.Product
	curProd.GetQueryParams(c)
	curProd.GetQueryId(c)
	return services.ProductShow(curProd).GetError(c)
}

func insertProduct(c *fiber.Ctx) error {
	var curProd tables.Product
	curProd.GetQueryParams(c)
	return services.ProductInsert(curProd).GetError(c)
}

func deleteProduct(c *fiber.Ctx) error {
	var curProd tables.Product
	curProd.GetQueryId(c)
	return services.ProductDelete(curProd).GetError(c)
}

func updateProduct(c *fiber.Ctx) error {
	var curProd tables.Product
	curProd.GetQueryParams(c)
	curProd.GetQueryId(c)
	return services.ProductUpdate(curProd).GetError(c)
}
