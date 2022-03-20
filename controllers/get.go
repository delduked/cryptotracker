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
	var trade types.TradeRequest

	// get all keyed fields in redis
	length, err := config.Rdb.Keys(config.RedisCtx, "*").Result()
	if err != nil {
		return tradesSaved, err
	}

	for _, j := range length {
		err := config.Rdb.HGetAll(config.RedisCtx, j).Scan(&trade)
		if err != nil {
			return tradesSaved, err
		}
		tradeSaved := types.TradeSaved{
			Key:              j,
			Date:             trade.Date,
			Type:             trade.Type,
			ReceivedQuantity: trade.ReceivedQuantity,
			ReceivedCurrency: trade.ReceivedCurrency,
			SentQuantity:     trade.SentQuantity,
			SentCurrency:     trade.SentCurrency,
			FeeAmount:        trade.FeeAmount,
			FeeCurrency:      trade.FeeCurrency,
			Tag:              trade.Tag,
		}
		// append the keyed field to the array of trades to be sent to the client
		tradesSaved = append(tradesSaved, tradeSaved)

	}
	return tradesSaved, err
}

func GetKeyedTrade(key string) (types.TradeSaved, error, int) {
	var trade types.TradeRequest
	var tradeWithKey types.TradeSaved
	var Error error
	length, err := config.Rdb.HGetAll(config.RedisCtx, key).Result()
	if err == redis.Nil {
		return tradeWithKey, Error, len(length)
	} else if err != nil {
		return tradeWithKey, Error, len(length)
	} else if len(length) == 0 {
		return tradeWithKey, Error, len(length)
	}

	err = config.Rdb.HGetAll(config.RedisCtx, key).Scan(&trade)
	if err != nil {
		return tradeWithKey, err, len(length)
	}

	tradeWithKey = types.TradeSaved{
		Key:              key,
		Date:             trade.Date,
		Type:             trade.Type,
		ReceivedQuantity: trade.ReceivedQuantity,
		ReceivedCurrency: trade.ReceivedCurrency,
		SentQuantity:     trade.SentQuantity,
		SentCurrency:     trade.SentCurrency,
		FeeAmount:        trade.FeeAmount,
		FeeCurrency:      trade.FeeCurrency,
		Tag:              trade.Tag,
	}

	return tradeWithKey, nil, len(length)

}
