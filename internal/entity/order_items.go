package entity

import "fmt"

type OrderItems struct {
	Item
	RID         string `json:"rid" example:"rid-4ad61f67-9281-4997-a411-c3e21fe5823c"`
	TrackNumber string `json:"track_number" example:"WB-4ad61f67-9281-4997-a411-c3e21fe5823c"`
	Sale        int    `json:"sale" example:"15"`
	Size        string `json:"size" example:"1"`
	TotalPrice  int    `json:"total_price" example:"5000"`
	Status      int    `json:"status" example:"202"`
}

func (oi *OrderItems) Validate(i int, trackNumber string) error {
	if oi.RID == "" {
		return fmt.Errorf("items[%d] rid is required", i)
	}

	if oi.TrackNumber == "" {
		return fmt.Errorf("items[%d] track_number is required", i)
	}
	if oi.TrackNumber != trackNumber {
		return fmt.Errorf("items[%d] track_number must be equal to order track_number", i)
	}

	if oi.Size == "" {
		return fmt.Errorf("items[%d] size is required", i)
	}

	if oi.TotalPrice <= 0 {
		return fmt.Errorf("items[%d] total_price must be positive", i)
	}

	if oi.Status == 0 {
		return fmt.Errorf("items[%d] status is required", i)
	}
	return nil
}
