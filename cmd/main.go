package main

import (
	"log"

	"github.com/SahithyTumma/golangApplication/database"
	"github.com/SahithyTumma/golangApplication/handlers"
	"github.com/SahithyTumma/golangApplication/pulsar"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	pulsar.InitializePulsarClient()
	defer pulsar.ClosePulsarClient()

	app := fiber.New()

	// routes
	app.Get("/", handlers.ListFacts)
	app.Post("/fact", handlers.CreateFact)
	app.Get("/fact/:id", handlers.ShowFact)
	app.Patch("/fact/:id", handlers.UpdateFact)
	app.Delete("/fact/:id", handlers.DeleteFact)

	go pulsar.ConsumeMessages()

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
