package main

import (
	"github.com/SahithyTumma/docker/database"
	"github.com/SahithyTumma/docker/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)

	app.Get("/fact/:id", handlers.ShowFact)

	app.Patch("/fact/:id", handlers.UpdateFact)

	app.Delete("/fact/:id", handlers.DeleteFact)

	app.Listen(":3000")
}
