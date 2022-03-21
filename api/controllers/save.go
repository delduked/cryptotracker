package controllers

import (
	"api/config"
	"api/types"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

func Save(value *types.TradeRequest) (types.TradeSaved, error) {
	key := (uuid.New()).String()
	var tradeSaved types.TradeSaved

	// Set some fields.
	_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
		rdb.HSet(config.RedisCtx, key, "Date", value.Date)
		rdb.HSet(config.RedisCtx, key, "Type", value.Type)
		rdb.HSet(config.RedisCtx, key, "ReceivedQuantity", value.ReceivedQuantity)
		rdb.HSet(config.RedisCtx, key, "ReceivedCurrency", value.ReceivedCurrency)
		rdb.HSet(config.RedisCtx, key, "SentQuantity", value.SentQuantity)
		rdb.HSet(config.RedisCtx, key, "SentCurrency", value.SentCurrency)
		rdb.HSet(config.RedisCtx, key, "FeeAmount", value.FeeAmount)
		rdb.HSet(config.RedisCtx, key, "FeeCurrency", value.FeeCurrency)
		rdb.HSet(config.RedisCtx, key, "Tag", value.Tag)
		return nil
	})

	if err != nil {
		return tradeSaved, err
	}

	tradeSaved = types.TradeSaved{
		Key:              key,
		Date:             value.Date,
		Type:             value.Type,
		ReceivedQuantity: value.ReceivedQuantity,
		ReceivedCurrency: value.ReceivedCurrency,
		SentQuantity:     value.SentQuantity,
		SentCurrency:     value.SentCurrency,
		FeeAmount:        value.FeeAmount,
		FeeCurrency:      value.FeeCurrency,
		Tag:              value.Tag,
	}

	return tradeSaved, nil

}
func SaveMany(value []types.TradeRequest) error {
	var err error
	var key string

	// range of array values from request body
	for _, j := range value {

		// create new key for each iteration of the body's array
		key = (uuid.New()).String()
		_, err := config.Rdb.Pipelined(config.RedisCtx, func(rdb redis.Pipeliner) error {
			rdb.HSet(config.RedisCtx, key, "Date", j.Date)
			rdb.HSet(config.RedisCtx, key, "Type", j.Type)
			rdb.HSet(config.RedisCtx, key, "ReceivedQuantity", j.ReceivedQuantity)
			rdb.HSet(config.RedisCtx, key, "ReceivedCurrency", j.ReceivedCurrency)
			rdb.HSet(config.RedisCtx, key, "SentQuantity", j.SentQuantity)
			rdb.HSet(config.RedisCtx, key, "SentCurrency", j.SentCurrency)
			rdb.HSet(config.RedisCtx, key, "FeeAmount", j.FeeAmount)
			rdb.HSet(config.RedisCtx, key, "FeeCurrency", j.FeeCurrency)
			rdb.HSet(config.RedisCtx, key, "Tag", j.Tag)
			return nil
		})
		if err != nil {
			return err
		}

	}
	return err
}
