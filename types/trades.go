package types

type TradeRequest struct {
	Date             string `json:"Date"`
	Type             string `json:"Type"`
	ReceivedQuantity int    `json:"ReceivedQuantity"`
	ReceivedCurrency int    `json:"ReceivedCurrency"`
	SentQuantity     int    `json:"SentQuantity"`
	SentCurrency     string `json:"SentCurrency"`
	FeeAmount        int    `json:"FeeAmount"`
	FeeCurrency      int    `json:"FeeCurrency"`
	Tag              string `json:"Tag"`
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
	Key              string `json:"Key"`
	Date             string `json:"Date"`
	Type             string `json:"Type"`
	ReceivedQuantity int    `json:"ReceivedQuantity"`
	ReceivedCurrency int    `json:"ReceivedCurrency"`
	SentQuantity     int    `json:"SentQuantity"`
	SentCurrency     string `json:"SentCurrency"`
	FeeAmount        int    `json:"FeeAmount"`
	FeeCurrency      int    `json:"FeeCurrency"`
	Tag              string `json:"Tag"`
}
