package main

import (
	"cryptoTracker/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println(c)
		return c.Next()
	})

	app.Get("/", services.GetTrades)
	app.Get("/:key", services.GetTrade)
	app.Get("/filter/:type/:receivedcurrency", services.GetTradeType)

	app.Post("/", services.SaveTrades)
	app.Put("/", services.SaveTrade)

	app.Delete("/", services.DeleteAllTrades)

	stats := app.Group("/stats")
	stats.Get("/dashboard", monitor.New())
	stats.Get("/healthcheck", services.HealthCheck)

	app.Listen(":8080")

}
