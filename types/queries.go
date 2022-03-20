package types

type TradeFilter struct {
	Type             string `query:"Type"`
	ReceivedCurrency string `query:"ReceivedCurrency"`
}
