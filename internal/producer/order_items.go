package producer

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"orderService/internal/entity"
	"time"
)

func (p *Producer) generateOrderItems(ctx context.Context, trackNumber string) ([]entity.OrderItems, error) {
	var orderItems []entity.OrderItems
	itemCount := rand.Intn(5) + 1
	for i := 0; i < itemCount; i++ {
		itemID := int64(rand.Intn(500) + 1)
		item, err := p.pgRepo.GetItemByItemID(ctx, itemID)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				return nil, err
			}
			return nil, fmt.Errorf("generateItems p.pgRepo.GetItemByItemID: %w", err)
		}
		sale := rand.Intn(30) + 1
		RID := fmt.Sprintf("rid-%d", time.Now().UnixNano())
		size := fmt.Sprintf("%d", rand.Intn(5)+1)
		discount := float32(100-sale) / 100.0
		totalPrice := int(float32(item.Price) * discount)
		val := entity.OrderItems{
			Item:        *item,
			RID:         RID,
			TrackNumber: trackNumber,
			Sale:        sale,
			Size:        size,
			TotalPrice:  totalPrice,
			Status:      202,
		}
		orderItems = append(orderItems, val)
	}
	return orderItems, nil
}
