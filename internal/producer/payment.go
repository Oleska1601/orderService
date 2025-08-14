package producer

import (
	"fmt"
	"math/rand"
	"orderService/internal/entity"
	"time"
)

func countGoodsTotal(items []entity.OrderItems) int {
	summ := 0
	for _, item := range items {
		summ += item.TotalPrice
	}
	return summ
}

func generatePayment(OrderUID string, items []entity.OrderItems) entity.Payment {
	transaction := fmt.Sprintf("transaction-%d", time.Now().UnixNano())
	paymentDT := int64(rand.Intn(90000000) + 10000000)
	goodsTotal := countGoodsTotal(items)
	deliveryCost := rand.Intn(100) + 100
	amount := goodsTotal + deliveryCost
	return entity.Payment{
		OrderUID:     OrderUID,
		Transaction:  transaction,
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       amount,
		PaymentDT:    paymentDT,
		Bank:         "alpha",
		DeliveryCost: deliveryCost,
		GoodsTotal:   goodsTotal,
		CustomFee:    0,
	}

}
