package publisher

import (
	"analytic-reporting/internal/publisher/sale"
	"database/sql"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

func InitPublisherService(router fiber.Router, db *sql.DB, kafkaProducer sarama.SyncProducer) {
	// SALE Module
	sale.InitSaleModule(router, db, kafkaProducer)
}
