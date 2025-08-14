package producer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"math/rand"
	"orderService/internal/entity"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func (p *Producer) generateOrder(ctx context.Context) *entity.Order {
	locales := []string{"en", "ru"}
	deliveryServices := []string{"test_delivery_service1", "test_delivery_service2", "test_delivery_service3", "test_delivery_service4", "test_delivery_service5"}
	id := uuid.New().String()
	orderUID := fmt.Sprintf("order-%s", id)
	trackNumber := fmt.Sprintf("WB-%s", id)
	delivery := generateDelivery(orderUID)
	customerID := fmt.Sprintf("customer-%s", id)
	items, err := p.generateOrderItems(ctx, trackNumber)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil
		}
		p.logger.Error("generateOrder p.generateItems", slog.Any("error", err))
		return nil
	}
	payment := generatePayment(orderUID, items)
	locale := locales[rand.Intn(2)]
	deliveryService := deliveryServices[rand.Intn(5)]
	shardkey := rand.Intn(10) + 1
	smID := rand.Intn(100) + 1
	oofShard := strconv.Itoa(rand.Intn(10) + 1)
	order := entity.Order{
		OrderUID:          orderUID,
		TrackNumber:       trackNumber,
		Entry:             "WBIL",
		Delivery:          delivery,
		Payment:           payment,
		Items:             items,
		Locale:            locale,
		InternalSignature: "",
		CustomerID:        customerID,
		DeliveryService:   deliveryService,
		ShardKey:          shardkey,
		SmID:              smID,
		DateCreated:       time.Now(),
		OofShard:          oofShard,
	}
	return &order
}

func (p *Producer) generateOrdersChan(ctx context.Context) chan *entity.Order {
	ordersChan := make(chan *entity.Order)
	go func() {
		defer close(ordersChan)
		for {
			order := p.generateOrder(ctx)
			if order == nil {
				continue
			}
			select {
			case <-ctx.Done():
				return
			case ordersChan <- order:
				time.Sleep(5 * time.Second) // небольшая тестовая задержка

			}
		}

	}()
	return ordersChan
}
