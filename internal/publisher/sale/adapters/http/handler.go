package httpadapter

import (
	postgresadapter "analytic-reporting/internal/publisher/sale/adapters/postgres"
	"analytic-reporting/internal/publisher/sale/app"
	"analytic-reporting/internal/publisher/sale/domain"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service app.Service
}

func NewHandler(svc app.Service) *Handler {
	return &Handler{Service: svc}
}

func (h *Handler) CreateSale(service app.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody domain.CreateSales
		err := c.BodyParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(domain.SaleErrorResponse(err))
		}
		result, err := service.CreateSales(&requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(domain.SaleErrorResponse(err))
		}
		return c.JSON(domain.SalesCreatedSuccessResponse(result))
	}
}

func (h *Handler) GetSales(service app.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var requestBody postgresadapter.QueryParams
		err := c.QueryParser(&requestBody)
		if err != nil {
			c.Status(http.StatusBadRequest)
			return c.JSON(domain.SaleErrorResponse(err))
		}
		result, err := service.GetSales(requestBody)
		if err != nil {
			c.Status(http.StatusInternalServerError)
			return c.JSON(domain.SaleErrorResponse(err))
		}
		var mapToDto []domain.SaleDTO
		for _, sale := range result {
			mapToDto = append(mapToDto, domain.SaleDTO{
				ID:           sale.ID,
				ShipDate:     sale.ShipDate,
				ShipMode:     sale.ShipMode,
				CustomerName: sale.CustomerName,
				Qty:          sale.Qty,
				SaleAmount:   sale.SaleAmount,
				Discount:     sale.Discount,
				Profit:       sale.Profit,
				ProfitRatio:  sale.ProfitRatio,
				OrderID:      sale.OrderID,
				OrderDate:    sale.OrderDate,
				Address:      sale.Address,
				Product:      sale.ProductName,
				Segment:      sale.SegmentName,
			})
		}
		return c.JSON(domain.SalesSuccessResponse(&mapToDto))
	}
}
