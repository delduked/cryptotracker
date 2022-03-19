package controllers

import (
	"cryptoTracker/config"
	"cryptoTracker/types"

	"github.com/go-redis/redis/v8"
)

func GetAll() ([]types.TradeSaved, error) {
	// create array of trades to be sent to the client
	tradesSaved := []types.TradeSaved{}

	// create each field in the array of trades to be saved
	var keyedTrade types.TradeSaved

	// get all keyed fields in redis
	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		return tradesSaved, err
	}

	for _, j := range length {
		err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&keyedTrade)
		if err != nil {
			return tradesSaved, err
		}
		tradeSaved := types.TradeSaved{
			Key:              j,
			Date:             keyedTrade.Date,
			Type:             keyedTrade.Type,
			ReceivedQuantity: keyedTrade.ReceivedQuantity,
			ReceivedCurrency: keyedTrade.ReceivedCurrency,
			SentQuantity:     keyedTrade.SentQuantity,
			SentCurrency:     keyedTrade.SentCurrency,
			FeeAmount:        keyedTrade.FeeAmount,
			FeeCurrency:      keyedTrade.FeeCurrency,
			Tag:              keyedTrade.Tag,
		}
		// append the keyed field to the array of trades to be sent to the client
		tradesSaved = append(tradesSaved, tradeSaved)

	}
	return tradesSaved, err
}

func GetKeyedTrade(key string) (types.TradeSaved, error, int) {
	var keyedTrade types.TradeSaved
	var Error error
	length, err := config.Rdb.HGetAll(config.RedisCtx, key).Result()
	if err == redis.Nil {
		return keyedTrade, Error, len(length)
	} else if err != nil {
		return keyedTrade, Error, len(length)
	} else if len(length) == 0 {
		return keyedTrade, Error, len(length)
	}

	err = config.Rdb.HGetAll(config.RedisCtx, key).Scan(&keyedTrade)
	if err != nil {
		return keyedTrade, err, len(length)
	}
	return keyedTrade, nil, len(length)

}
