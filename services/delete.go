package services

import (
	"cryptoTracker/controllers"
	"cryptoTracker/handlers"
	"cryptoTracker/types"

	"github.com/gofiber/fiber/v2"
)

func DeleteAllTrades(c *fiber.Ctx) error {

	// delete all the trades in redis
	err := controllers.DeleteAll()
	if err != nil {
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	// construct the response
	res := types.Response{
		Status: fiber.StatusOK,
		Error:  nil,
	}
	return handlers.Response(res, c)

}
