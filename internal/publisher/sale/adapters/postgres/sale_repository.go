package postgresadapter

import (
	postgresadapter "analytic-reporting/internal/publisher/product/adapters/postgres"
	product "analytic-reporting/internal/publisher/product/domain"
	"analytic-reporting/internal/publisher/sale/domain"
	log "analytic-reporting/internal/publisher/shared/domain"
	"analytic-reporting/model"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type QueryParams struct {
	OrderDate *string `query:"order_date"`
}

type Repository interface {
	CreateSales(sale *domain.CreateSales) (int32, error)
	GetSales(queryParams QueryParams) ([]model.SaleWithDetail, error)
}

type repository struct {
	db                *sql.DB
	productRepository postgresadapter.Repository
}

func NewRepo(db *sql.DB, productRepository postgresadapter.Repository) Repository {
	return &repository{
		db:                db,
		productRepository: productRepository,
	}
}

func (r *repository) CreateSales(sale *domain.CreateSales) (int32, error) {
	locationUUID, err := uuid.Parse(sale.LocationID)
	if err != nil {
		fmt.Println("Invalid LocationID UUID:", err)
		return -1, err
	}

	segmentUUID, err := uuid.Parse(sale.SegmentID)
	if err != nil {
		fmt.Println("Invalid Segment UUID:", err)
		return -1, err
	}

	var productIDs []string
	for i := range len(sale.ProductDetails) {
		productIDs = append(productIDs, sale.ProductDetails[i].ProductID)
	}

	products, err := r.productRepository.GetProducts(
		&postgresadapter.QueryParams{
			IDs: productIDs,
		},
	)
	if err != nil {
		fmt.Println("Invalid Product:", err)
		return -1, err
	}

	productMap := make(map[string]product.Product)
	for i := range len(*products) {
		p := (*products)[i]
		productMap[p.ID.String()] = p
	}

	var saleEntities []domain.Sale

	for i, details := range sale.ProductDetails {
		productDetails := productMap[details.ProductID]

		pricePerProduct := sale.ProductDetails[i].SalesAmount /
			float32(sale.ProductDetails[i].Qty)

		basePrice := productDetails.BasePrice
		profit := pricePerProduct - basePrice
		profitRatio := profit / basePrice

		productUUID, err := uuid.Parse(sale.ProductDetails[i].ProductID)
		if err != nil {
			fmt.Println("Invalid Segment UUID:", err)
			return -1, err
		}

		logData := log.Log{
			CreatedBy: "System",
			CreatedAt: time.Now(),
			UpdatedBy: "System",
			UpdatedAt: time.Now(),
		}

		saleEntities = append(saleEntities, domain.Sale{
			ID:             uuid.New(),
			ShipDate:       sale.ShipDate,
			ShipMode:       sale.ShipMode,
			CustomerName:   sale.CustomerName,
			Qty:            sale.ProductDetails[i].Qty,
			SaleAmount:     sale.ProductDetails[i].SalesAmount,
			Discount:       sale.ProductDetails[i].Discount,
			Profit:         profit,
			ProfitRatio:    profitRatio,
			NumberOfRecord: int32(i + 1),
			OrderID:        sale.OrderID,
			OrderDate:      sale.OrderDate,
			LocationID:     locationUUID,
			ProductID:      productUUID,
			SegmentID:      segmentUUID,
			Log:            logData,
		})
	}

	query := `
		INSERT INTO sales (
			id, ship_date, ship_mode, customer_name, quantity, sales_amount, discount,
			profit, profit_ratio, number_of_record, order_id, order_date,
			location_id, product_id, segment_id,
			created_by, created_at, updated_by, updated_at
		) VALUES 
	`

	var (
		args         []interface{}
		placeholders []string
	)

	for i, sale := range saleEntities {
		offset := i * 19 // 19 fields per row

		placeholders = append(placeholders, fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d,$%d)",
			offset+1, offset+2, offset+3, offset+4, offset+5, offset+6, offset+7,
			offset+8, offset+9, offset+10, offset+11, offset+12, offset+13, offset+14, offset+15,
			offset+16, offset+17, offset+18, offset+19,
		))

		args = append(args,
			sale.ID,
			sale.ShipDate,
			sale.ShipMode,
			sale.CustomerName,
			sale.Qty,
			sale.SaleAmount,
			sale.Discount,
			sale.Profit,
			sale.ProfitRatio,
			sale.NumberOfRecord,
			sale.OrderID,
			sale.OrderDate,
			sale.LocationID,
			sale.ProductID,
			sale.SegmentID,
			sale.Log.CreatedBy,
			sale.Log.CreatedAt,
			sale.Log.UpdatedBy,
			sale.Log.UpdatedAt,
		)
	}

	// Combine full query
	query += strings.Join(placeholders, ", ")

	// Execute
	rows, err := r.db.Exec(query, args...)
	if err != nil {
		return -1, fmt.Errorf("bulk insert failed: %w", err)
	}

	createdRows, err := rows.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("something wrong: %w", err)
	}

	return int32(createdRows), nil
}

func (r *repository) GetSales(queryParams QueryParams) ([]model.SaleWithDetail, error) {
	baseQuery := `select
			s.id,
			s.ship_date,
			s.ship_mode,
			s.customer_name,
			s.quantity,
			s.sales_amount,
			s.discount,
			s.profit,
			s.profit_ratio,
			s.number_of_record,
			s.order_id,
			s.order_date,
			s.location_id,
			s.product_id,
			s.segment_id,
			concat(l.city, ', ', l.state, ', ', l.postal_code) as address,
			p.product_name ,
			sg.segment_name 
		from
			sales s
		left join locations l on
			l.id = s.location_id
		left join products p on
			p.id = s.product_id
		left join segments sg on
			sg.id = s.segment_id
		where
			1 = 1
	`
	var args []any
	argIndex := 1

	if queryParams.OrderDate != nil && *queryParams.OrderDate != "" {
		baseQuery += ` and date(order_date) = $` + strconv.Itoa(argIndex)
		args = append(args, queryParams.OrderDate)
		argIndex++
	}

	rows, err := r.db.Query(baseQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sales []model.SaleWithDetail
	for rows.Next() {
		var sale model.SaleWithDetail

		if err := rows.Scan(
			&sale.ID,
			&sale.ShipDate,
			&sale.ShipMode,
			&sale.CustomerName,
			&sale.Qty,
			&sale.SaleAmount,
			&sale.Discount,
			&sale.Profit,
			&sale.ProfitRatio,
			&sale.NumberOfRecord,
			&sale.OrderID,
			&sale.OrderDate,
			&sale.LocationID,
			&sale.ProductID,
			&sale.SegmentID,
			&sale.Address,
			&sale.ProductName,
			&sale.SegmentName,
		); err != nil {
			return nil, err
		}
		sales = append(sales, sale)
	}

	return sales, nil
}
