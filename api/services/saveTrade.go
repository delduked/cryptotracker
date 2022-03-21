package services

import (
	"api/controllers"
	"api/handlers"
	"api/types"

	"github.com/gofiber/fiber/v2"
)

func SaveTrade(c *fiber.Ctx) error {
	// only accept POST requests
	c.Accepts("application/json")

	// check if the request body matches the requires body type
	body := new(types.TradeRequest)
	if err := c.BodyParser(body); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	// save the trade in redis
	savedTrade, err := controllers.Save(body)
	if err != nil {
		res := types.Response{
			Status: fiber.StatusInternalServerError,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	// construct the response
	res := types.NewTradeResponse{
		Status: fiber.StatusOK,
		Error:  nil,
		Trade:  savedTrade,
	}

	return handlers.NewTradeResponse(res, c)
}

func SaveTrades(c *fiber.Ctx) error {
	// only accept POST requests
	c.Accepts("application/json")

	// check if the request body matches the requires body type
	body := new([]types.TradeRequest)
	if err := c.BodyParser(body); err != nil {
		res := types.Response{
			Status: fiber.StatusBadRequest,
			Error:  err,
		}
		return handlers.Response(res, c)
	}

	// save the trades in redis
	err := controllers.SaveMany(*body)
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
