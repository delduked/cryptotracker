package controllers

import (
	"cryptoTracker/types"
)

func GetTradeType(property *types.TradeFilter) ([]types.TradeSaved, error) {
	var deposits []types.TradeSaved

	allTrades, err := GetAll()
	if err != nil {
		return deposits, err
	}

	for _, trade := range allTrades {

		// check which filters user is requesting
		typeErr := ValidateTypeProperty(trade.Type)
		receivedCurrencyErr := ValidateReceivedCurrency(trade.ReceivedCurrency)

		// response to user with both filters as requested even if the filter is empty
		if typeErr == nil && receivedCurrencyErr == nil {
			if trade.Type == property.Type && trade.ReceivedCurrency == property.ReceivedCurrency {

				deposit := types.TradeSaved{
					Key:              trade.Key,
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

				deposits = append(deposits, deposit)
			}
		}

		// response to user with the type filter even if the received currency filter is anything
		if typeErr == nil && receivedCurrencyErr != nil {
			if trade.Type == property.Type {

				deposit := types.TradeSaved{
					Key:              trade.Key,
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

				deposits = append(deposits, deposit)
			}
		}

		// response to user with the received currency filter even if the type filter is anything
		if typeErr != nil && receivedCurrencyErr == nil {
			if trade.ReceivedCurrency == property.ReceivedCurrency {

				deposit := types.TradeSaved{
					Key:              trade.Key,
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

				deposits = append(deposits, deposit)
			}
		}

		// response to user when no query paramater is provided, essentially respond with everything
		if typeErr != nil && receivedCurrencyErr != nil {
			deposit := types.TradeSaved{
				Key:              trade.Key,
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

			deposits = append(deposits, deposit)
		}
	}
	// response will be blank if no trades match the query
	return deposits, nil
}
