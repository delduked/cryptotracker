package handlers

import (
	"cryptoTracker/types"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func Response(res types.Response, c *fiber.Ctx) error {
	writer := c.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func NewTradeResponse(res types.NewTradeResponse, c *fiber.Ctx) error {
	writer := c.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func TradeResponse(res types.TradeResponse, c *fiber.Ctx) error {
	writer := c.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
func AllTradesResponse(res types.AllTradesResponse, c *fiber.Ctx) error {
	writer := c.Type("json", "utf-8").Response().BodyWriter()
	return json.NewEncoder(writer).Encode(res)
}
