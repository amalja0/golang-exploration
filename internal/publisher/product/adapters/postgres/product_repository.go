package postgresadapter

import (
	"analytic-reporting/internal/publisher/product/domain"
	"database/sql"
	"fmt"

	"github.com/lib/pq"
)

type QueryParams struct {
	IDs   []string
	Names []string
}

type Repository interface {
	GetProducts(params *QueryParams) (*[]domain.Product, error)
}

type repository struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetProducts(params *QueryParams) (*[]domain.Product, error) {
	baseQuery := `select * from products p where 1=1`
	var args []any
	argIndex := 1

	if len(params.IDs) > 0 {
		baseQuery += fmt.Sprintf(" and p.id = ANY ($%d)", argIndex)
		args = append(args, pq.Array(params.IDs))
		argIndex++
	}

	if len(params.Names) > 0 {
		baseQuery += fmt.Sprintf(" and p.name = ANY ($%d)", argIndex)
		args = append(args, pq.Array(params.Names))
		argIndex++
	}

	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product

		if err := rows.Scan(
			&product.ID,
			&product.ProductName,
			&product.Manufacturer,
			&product.BasePrice,
			&product.CreatedBy,
			&product.UpdatedBy,
			&product.DeletedBy,
			&product.CreatedAt,
			&product.UpdatedAt,
			&product.CategoryID,
			&product.SubCategoryID,
		); err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return &products, nil
}
