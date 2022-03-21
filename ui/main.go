package main

import (
	"ui/middleware"
	"ui/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Use(cors.New())

	ui := app.Group("/", middleware.Log)
	ui.Get("/", routes.Index)
	ui.Static("/assets", "./assets")

	app.Listen(":8080")

}
