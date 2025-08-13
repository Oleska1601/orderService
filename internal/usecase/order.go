package usecase

import (
	"context"
	"errors"
	"log/slog"
	"orderService/internal/entity"
)

func (u *Usecase) GetOrderByOrderUID(ctx context.Context, orderUID string) (*entity.Order, error) {
	order, ok := u.cache.Get(orderUID)
	if ok {
		return order, nil
	}
	order, err := u.pgRepo.GetOrderByOrderUID(ctx, orderUID)
	if order == nil {
		if err != nil {
			u.logger.Error("GetOrderByOrderUID u.pgRepo.GetOrderByOrderUID", slog.Any("error", err))
			return nil, errors.New("error of getting order") //internal server error
		}
		u.logger.Error("GetOrderByOrderUID u.pgRepo.GetOrderByOrderUID", slog.Any("error", "no order with current orderUID"), slog.String("orderUID", orderUID))
		return nil, errors.New("not found")

	}
	u.cache.Set(orderUID, *order)
	return order, nil
}
