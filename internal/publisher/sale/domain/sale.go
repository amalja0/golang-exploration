package domain

import (
	"analytic-reporting/internal/publisher/shared/domain"
	"time"

	"github.com/google/uuid"
)

type Sale struct {
	ID             uuid.UUID `db:"id" json:"id"`
	ShipDate       time.Time `db:"ship_date" json:"ship_date"`
	ShipMode       string    `db:"ship_mode" json:"ship_mode"`
	CustomerName   string    `db:"customer_name" json:"customer_name"`
	Qty            int32     `db:"quantity" json:"quantity"`
	SaleAmount     float32   `db:"sale_amount" json:"sale_amount"`
	Discount       float32   `db:"discount" json:"discount"`
	Profit         float32   `db:"profit" json:"profit"`
	ProfitRatio    float32   `db:"profit_ratio" json:"profit_ratio"`
	NumberOfRecord int32     `db:"number_of_record" json:"number_of_record"`
	OrderID        string    `db:"order_id" json:"order_id"`
	OrderDate      time.Time `db:"order_date" json:"order_date"`
	LocationID     uuid.UUID `db:"location_id" json:"location_id"`
	ProductID      uuid.UUID `db:"product_id" json:"product_id"`
	SegmentID      uuid.UUID `db:"segment_id" json:"segment_id"`
	domain.Log
}

type ProductDetail struct {
	ProductID   string  `json:"product_id"`
	Discount    float32 `json:"discount"`
	Qty         int32   `json:"quantity"`
	SalesAmount float32 `json:"sales_amount"`
}

type CreateSales struct {
	ShipDate       time.Time       `json:"ship_date"`
	ShipMode       string          `json:"ship_mode"`
	CustomerName   string          `json:"customer_name"`
	OrderID        string          `json:"order_id"`
	OrderDate      time.Time       `json:"order_date"`
	LocationID     string          `json:"location_id"`
	ProductDetails []ProductDetail `json:"order_details"`
	SegmentID      string          `json:"segment_id"`
}
