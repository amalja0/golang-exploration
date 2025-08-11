package model

import (
	"time"

	"github.com/google/uuid"
)

type Sale struct {
	ID             uuid.UUID `json:"id"`
	ShipDate       time.Time `json:"ship_date"`
	ShipMode       string    `json:"ship_mode"`
	CustomerName   string    `json:"customer_name"`
	Qty            int32     `json:"quantity"`
	SaleAmount     float32   `json:"sale_amount"`
	Discount       float32   `json:"discount"`
	Profit         float32   `json:"profit"`
	ProfitRatio    float32   `json:"profit_ratio"`
	NumberOfRecord int32     `json:"number_of_record"`
	OrderID        string    `json:"order_id"`
	OrderDate      time.Time `json:"order_date"`
	LocationID     uuid.UUID `json:"location_id"`
	ProductID      uuid.UUID `json:"product_id"`
	SegmentID      uuid.UUID `json:"segment_id"`
}

type SaleWithDetail struct {
	Sale
	Address     string `json:"address"`
	ProductName string `json:"product_name"`
	SegmentName string `json:"segment_name"`
}
