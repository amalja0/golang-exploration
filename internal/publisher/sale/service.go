package sale

import (
	"analytic-reporting/internal/publisher/product"
	httpadapter "analytic-reporting/internal/publisher/sale/adapters/http"
	postgresadapter "analytic-reporting/internal/publisher/sale/adapters/postgres"
	"analytic-reporting/internal/publisher/sale/app"
	"database/sql"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

func InitSaleModule(router fiber.Router, db *sql.DB, kafkaProducer sarama.SyncProducer) {
	productRepo := product.Init(db)
	repo := postgresadapter.NewRepo(db, productRepo)
	svc := app.NewService(repo, kafkaProducer)

	httpadapter.SaleRouter(router, svc)
}
