package entity

type Delivery struct {
	OrderUID string `json:"order_uid,omitempty"`
	Name     string `json:"name" example:"Иван Иванов"`
	Phone    string `json:"phone" example:"+7 910 000-0001"`
	Zip      string `json:"zip" example:"101000"`
	City     string `json:"city" example:"Москва"`
	Address  string `json:"address" example:"ул. Ленина, д. 1, кв. 1"`
	Region   string `json:"region" example:"Московская область"`
	Email    string `json:"email" example:"a@example.com"`
}
