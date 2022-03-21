package services

import (
	"api/handlers"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

func HealthCheck(c *fiber.Ctx) error {

	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}

	return handlers.Response(res, c)

}
