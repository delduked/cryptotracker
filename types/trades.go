package types

type TradeRequest struct {
	Date             string  `redis:"Date" json:"Date"`
	Type             string  `redis:"Type" json:"Type"`
	ReceivedQuantity float64 `redis:"ReceivedQuantity" json:"ReceivedQuantity"`
	ReceivedCurrency string  `redis:"ReceivedCurrency" json:"ReceivedCurrency"`
	SentQuantity     float64 `redis:"SentQuantity" json:"SentQuantity"`
	SentCurrency     string  `redis:"SentCurrency" json:"SentCurrency"`
	FeeAmount        float64 `redis:"FeeAmount" json:"FeeAmount"`
	FeeCurrency      string  `redis:"FeeCurrency" json:"FeeCurrency"`
	Tag              string  `redis:"Tag" json:"Tag"`
}
type NewTradeResponse struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
	Trade  TradeSaved
}
type AllTradesResponse struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
	Trades []TradeSaved
}
type TradeResponse struct {
	Status int   `json:"Status"`
	Error  error `json:"Error"`
	Trade  TradeSaved
}
type TradeSaved struct {
	Key              string  `redis:"Key" json:"Key"`
	Date             string  `redis:"Date" json:"Date"`
	Type             string  `redis:"Type" json:"Type"`
	ReceivedQuantity float64 `redis:"ReceivedQuantity" json:"ReceivedQuantity"`
	ReceivedCurrency string  `redis:"ReceivedCurrency" json:"ReceivedCurrency"`
	SentQuantity     float64 `redis:"SentQuantity" json:"SentQuantity"`
	SentCurrency     string  `redis:"SentCurrency" json:"SentCurrency"`
	FeeAmount        float64 `redis:"FeeAmount" json:"FeeAmount"`
	FeeCurrency      string  `redis:"FeeCurrency" json:"FeeCurrency"`
	Tag              string  `redis:"Tag" json:"Tag"`
}
