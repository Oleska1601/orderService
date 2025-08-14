package entity

import (
	"errors"
	"strings"
)

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

func (d *Delivery) Validate() error {
	// Проверка
	if d.Name == "" {
		return errors.New("delivery name is required")
	}

	if d.Phone == "" {
		return errors.New("delivery phone is required")
	}

	if d.Zip == "" {
		return errors.New("delivery zip is required")
	}

	if d.City == "" {
		return errors.New("delivery city is required")
	}

	if d.Address == "" {
		return errors.New("delivery address is required")
	}

	if d.Region == "" {
		return errors.New("delivery region is required")
	}

	if d.Email == "" {
		return errors.New("delivery email is required")
	} else if !strings.Contains(d.Email, "@") {
		return errors.New("delivery email must be a valid email address")
	}
	return nil
}
