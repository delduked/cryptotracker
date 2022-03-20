package services

import (
	"cryptoTracker/controllers"
	"cryptoTracker/handlers"
	"cryptoTracker/types"

	"github.com/gofiber/fiber/v2"
)

func GetTrades(c *fiber.Ctx) error {
	tradesSaved, err := controllers.GetAll()
	if err != nil {
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := types.AllTradesResponse{
		Status: fiber.StatusOK,
		Error:  err,
		Trades: tradesSaved,
	}

	return handlers.AllTradesResponse(res, c)
}

func GetTrade(c *fiber.Ctx) error {
	key := c.Params("key")

	trade, err, length := controllers.GetKeyedTrade(key)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	} else if length < 1 {
		res := types.Response{
			Status: fiber.StatusNotFound,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	res := types.TradeResponse{
		Status: fiber.StatusOK,
		Error:  err,
		Trade:  trade,
	}
	return handlers.TradeResponse(res, c)

}
