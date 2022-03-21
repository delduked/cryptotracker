package main

import (
	"api/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	//app.Get("/swagger/*", swagger.Handler)

	app.Get("/dashboard", monitor.New())
	app.Get("/healthcheck", services.HealthCheck)

	db := app.Group("/db")

	db.Get("/", services.GetTrades)
	db.Get("/:key", services.GetTrade)
	db.Post("/", services.SaveTrades)
	db.Put("/", services.SaveTrade)
        db.Delete("/", services.DeleteAllTrades)

	app.Listen(":8080")

}
