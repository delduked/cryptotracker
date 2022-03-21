package routes

import (
	"ui/config.go"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	// Render index template
	return c.Render("index", fiber.Map{
		"API_HOSTNAME": config.Hostname,
	})
}
