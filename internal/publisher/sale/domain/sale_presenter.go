package domain

import (
	"analytic-reporting/model"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type SaleDTO struct {
	ID           uuid.UUID `json:"id"`
	ShipDate     time.Time `json:"ship_date"`
	ShipMode     string    `json:"ship_mode"`
	CustomerName string    `json:"customer_name"`
	Qty          int32     `json:"quantity"`
	SaleAmount   float32   `json:"sale_amount"`
	Discount     float32   `json:"discount"`
	Profit       float32   `json:"profit"`
	ProfitRatio  float32   `json:"profit_ratio"`
	OrderID      string    `json:"order_id"`
	OrderDate    time.Time `json:"order_date"`
	Address      string    `json:"address"`
	Product      string    `json:"product"`
	Segment      string    `json:"segment"`
}

func SaleSuccessResponse(data *model.SaleWithDetail) *fiber.Map {
	sale := SaleDTO{
		ID:           data.ID,
		ShipDate:     data.ShipDate,
		ShipMode:     data.ShipMode,
		CustomerName: data.CustomerName,
		Qty:          data.Qty,
		SaleAmount:   data.SaleAmount,
		Discount:     data.Discount,
		Profit:       data.Profit,
		ProfitRatio:  data.ProfitRatio,
		OrderID:      data.OrderID,
		OrderDate:    data.OrderDate,
		Address:      data.Address,
		Product:      data.ProductName,
		Segment:      data.SegmentName,
	}
	return &fiber.Map{
		"status": true,
		"data":   sale,
		"error":  nil,
	}
}

func SalesSuccessResponse(data *[]SaleDTO) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func SalesCreatedSuccessResponse(rowsCreated int32) *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   rowsCreated,
		"error":  nil,
	}
}

func SaleErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   "",
		"error":  err.Error(),
	}
}
