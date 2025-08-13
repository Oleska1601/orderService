package entity

type Payment struct {
	OrderUID     string  `json:"order_uid,omitempty"`
	Transaction  string  `json:"transaction" example:"transaction-4ad61f67-9281-4997-a411-c3e21fe5823c"`
	RequestID    string  `json:"request_id" example:""`
	Currency     string  `json:"currency" example:"USD"`
	Provider     string  `json:"provider" example:"wbpay"`
	Amount       float64 `json:"amount" example:"5300"`
	PaymentDT    int64   `json:"payment_dt" example:"1637907727"`
	Bank         string  `json:"bank" example:"alpha"`
	DeliveryCost float64 `json:"delivery_cost" example:"300"`
	GoodsTotal   int64   `json:"goods_total" example:"5000"`
	CustomFee    int64   `json:"custom_fee" example:"0"`
}
