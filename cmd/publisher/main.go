package main

import (
	"analytic-reporting/internal/publisher"
	"analytic-reporting/internal/publisher/dbhandler"
	kafkaConfig "analytic-reporting/internal/publisher/kafka"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db, err := dbhandler.ConnectPostgres()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	kafka, err := kafkaConfig.ConnectKafkaProducer()
	if err != nil {
		log.Fatal(err)
	}
	defer kafka.Close()

	// Fiber instance
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("Welcome to the trial of creating go rest!"))
	})
	api := app.Group("/api")

	publisher.InitPublisherService(api, db, kafka)

	log.Fatal(app.Listen(":8080"))
}
