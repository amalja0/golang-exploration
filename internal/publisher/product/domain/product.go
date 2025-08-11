package domain

import (
	"analytic-reporting/internal/publisher/shared/domain"

	"github.com/google/uuid"
)

type Product struct {
	ID            uuid.UUID `db:"id"`
	ProductName   string    `db:"product_name"`
	Manufacturer  string    `db:"manufacturer"`
	BasePrice     float32   `db:"base_price"`
	CategoryID    uuid.UUID `db:"category_id"`
	SubCategoryID uuid.UUID `db:"sub_category_id"`
	domain.Log
}
