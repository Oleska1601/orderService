package entity

type OrderItems struct {
	Item
	RID         string `json:"rid" example:"rid-4ad61f67-9281-4997-a411-c3e21fe5823c"`
	TrackNumber string `json:"track_number" example:"WB-4ad61f67-9281-4997-a411-c3e21fe5823c"`
	Sale        int32  `json:"sale" example:"15"`
	Size        string `json:"size" example:"1"`
	TotalPrice  int32  `json:"total_price" example:"5000"`
	Status      int    `json:"status" example:"202"`
}
