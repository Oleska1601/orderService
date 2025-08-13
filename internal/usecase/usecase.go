package usecase

import (
	"context"
	"fmt"
	"orderService/internal/cache"
	"orderService/internal/database/repo"
	"orderService/internal/entity"
	"orderService/pkg/logger"
)

type UsecaseInterface interface {
	GetOrderByOrderUID(context.Context, string) (*entity.Order, error)
}

type Usecase struct {
	cache  *cache.LRUCache
	pgRepo repo.PgRepoInterface
	logger logger.LoggerInterface
}

func NewUsecase(cache *cache.LRUCache, pgRepo repo.PgRepoInterface, logger logger.LoggerInterface) (*Usecase, error) {
	u := &Usecase{
		cache:  cache,
		pgRepo: pgRepo,
		logger: logger,
	}
	// Заполняем кэш при инициализации
	if err := u.fillCache(cache.Capacity); err != nil {
		return nil, fmt.Errorf("NewUsecase u.fillCache: %w", err)
	}
	u.logger.Info("initialise usecase successful")
	return u, nil
}

func (u *Usecase) fillCache(capacity int) error {
	orders, err := u.pgRepo.GetOrdersWithCapacity(context.Background(), capacity)
	if err != nil {
		return err
	}

	for _, order := range orders {
		u.cache.Set(order.OrderUID, order)
	}

	return nil
}
