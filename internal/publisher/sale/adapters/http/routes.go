package httpadapter

import (
	"analytic-reporting/internal/publisher/sale/app"

	"github.com/gofiber/fiber/v2"
)

func SaleRouter(app fiber.Router, service app.Service) {
	handlers := NewHandler(service)
	app.Post("/sale", handlers.CreateSale(service))
	app.Get("/sale", handlers.GetSales(service))
}
