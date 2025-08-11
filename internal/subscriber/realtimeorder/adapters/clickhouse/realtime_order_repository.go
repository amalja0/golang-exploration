package clickhouse

import (
	"analytic-reporting/internal/subscriber/realtimeorder/domain"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type RealtimeOrderRepository interface {
	InitDB() error
	CreateOrderRecord(record domain.OrderRecord) error
}

type realtimeOrderRepository struct {
	sqlFilePath string
	db          driver.Conn
}

func InitRealtimeRepository(sqlFilePath string, db *driver.Conn) RealtimeOrderRepository {
	return &realtimeOrderRepository{
		sqlFilePath: sqlFilePath,
		db:          *db,
	}
}

func (r *realtimeOrderRepository) InitDB() error {
	sqlBytes, err := os.ReadFile(r.sqlFilePath)
	if err != nil {
		return fmt.Errorf("failed to read SQL file: %w", err)
	}

	for _, stmt := range strings.Split(string(sqlBytes), ";") {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		if err := r.db.Exec(context.Background(), stmt); err != nil {
			return fmt.Errorf("failed to execute SQL statement: %q: %w", stmt, err)
		}
	}

	log.Println("✅ ClickHouse database initialized successfully.")
	return nil
}

func (r *realtimeOrderRepository) CreateOrderRecord(record domain.OrderRecord) error {
	query := `
		INSERT INTO realtime_order (
			id,
			sale_id,
			quantity,
			sale_amount,
			discount,
			profit,
			profit_ratio,
			order_id,
			order_date,
			location_id,
			product_id,
			segment_id,
			product_name,
			segment_name,
			created_at
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)
	`
	fmt.Println(record)
	err := r.db.Exec(
		context.Background(),
		query,
		record.Id,
		record.SaleId,
		record.Quantity,
		record.SaleAmount,
		record.Discount,
		record.Profit,
		record.ProfitRatio,
		record.OrderId,
		record.OrderDate,
		record.LocationId,
		record.ProductId,
		record.SegmentId,
		record.ProductName,
		record.SegmentName,
		record.CreatedAt,
	)
	if err != nil {
		fmt.Println("Fail to insert record:", err)
	}
	return err
}
