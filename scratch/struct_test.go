package scratch

type order struct {
	OrderNumber int64  `json:"order_number,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Note        string `json:"note,omitempty"`
}

type orderExchange struct {
	OrderNumber int64  `json:"order_number,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Note        string `json:"note,omitempty"`
}

type exchangeOrders struct {
	OrdersExchangeHist   []*ExchangeOrdersHist `json:"exchange_orders_hist,omitempty"`
	NumberExchange       int64                 `json:"exchange_number,omitempty"`
	ExchangeLimitReached bool                  `json:"exchange_limit_reached,omitempty"`
}

type exchangeOrdersHist struct {
	OrderID       int64 `json:"order_id,omitempty"`
	ParentOrderID int64 `json:"parent_order_id,omitempty"`
}

type orderInit struct {
	OrderNumber int64  `json:"order_number,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Note        string `json:"note,omitempty"`
}
