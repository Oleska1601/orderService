package entity

type Item struct {
	ItemID int64  `json:"item_id,omitempty"`
	ChrtID int64  `json:"chrt_id" example:"5066679"`
	Price  int    `json:"price" example:"4499"`
	Name   string `json:"name" example:"Test Product Name 1"`
	NmID   int64  `json:"nm_id" example:"6332706"`
	Brand  string `json:"brand" example:"Brand-W2"`
}
