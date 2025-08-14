package entity

import (
	"errors"
	"time"
)

type Order struct {
	OrderUID          string       `json:"order_uid" example:"order-4ad61f67-9281-4997-a411-c3e21fe5823c"`
	TrackNumber       string       `json:"track_number" example:"WB-4ad61f67-9281-4997-a411-c3e21fe5823c"`
	Entry             string       `json:"entry" example:"WBIL"`
	Delivery          Delivery     `json:"delivery"`
	Payment           Payment      `json:"payment"`
	Items             []OrderItems `json:"items"`
	Locale            string       `json:"locale" example:"en"`
	InternalSignature string       `json:"internal_signature" example:""`
	CustomerID        string       `json:"customer_id" example:"1"`
	DeliveryService   string       `json:"delivery_service" example:"test_delivery_service1"`
	ShardKey          int          `json:"shardkey" example:"9"`
	SmID              int          `json:"sm_id" example:"99"`
	DateCreated       time.Time    `json:"date_created" example:"2021-11-26T06:22:19Z"`
	OofShard          string       `json:"oof_shard" example:"1"`
}

func (o *Order) Validate() error {
	if o.OrderUID == "" {
		return errors.New("order_uid is required")
	}

	if o.TrackNumber == "" {
		return errors.New("track_number is required")
	}

	if o.Entry == "" {
		return errors.New("entry is required")
	}

	if o.Locale == "" {
		return errors.New("locale is required")
	}

	if o.CustomerID == "" {
		return errors.New("customer_id is required")
	}

	if o.ShardKey == 0 {
		return errors.New("shardkey is required")
	}

	if o.SmID == 0 {
		return errors.New("sm_id is required")
	}

	if o.DateCreated.IsZero() {
		return errors.New("date_created is required")
	}

	if o.OofShard == "" {
		return errors.New("oof_shard is required")
	}

	if err := o.Delivery.Validate(); err != nil {
		return err
	}

	if err := o.Payment.Validate(); err != nil {
		return err
	}

	if len(o.Items) == 0 {
		return errors.New("order must contain at least one item")
	}

	for i, item := range o.Items {
		if err := item.Validate(i, o.TrackNumber); err != nil {
			return err
		}
	}

	return nil
}
