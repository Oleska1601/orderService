package entity

import "errors"

type Payment struct {
	OrderUID     string `json:"order_uid,omitempty"`
	Transaction  string `json:"transaction" example:"transaction-4ad61f67-9281-4997-a411-c3e21fe5823c"`
	RequestID    string `json:"request_id" example:""`
	Currency     string `json:"currency" example:"USD"`
	Provider     string `json:"provider" example:"wbpay"`
	Amount       int    `json:"amount" example:"5300"`
	PaymentDT    int64  `json:"payment_dt" example:"1637907727"`
	Bank         string `json:"bank" example:"alpha"`
	DeliveryCost int    `json:"delivery_cost" example:"300"`
	GoodsTotal   int    `json:"goods_total" example:"5000"`
	CustomFee    int    `json:"custom_fee" example:"0"`
}

func (p *Payment) Validate() error {
	if p.Transaction == "" {
		return errors.New("payment transaction is required")
	}

	if p.Currency == "" {
		return errors.New("payment currency is required")
	}

	if p.Provider == "" {
		return errors.New("payment provider is required")
	}

	if p.Amount <= 0 {
		return errors.New("payment amount must be positive")
	}

	if p.PaymentDT <= 0 {
		return errors.New("payment payment_dt must be positive")
	}

	if p.Bank == "" {
		return errors.New("payment bank is required")
	}

	if p.GoodsTotal <= 0 {
		return errors.New("payment goods_total must be positive")
	}
	return nil
}
