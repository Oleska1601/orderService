package producer

import (
	"context"
	"fmt"
	"math/rand"
	"orderService/internal/entity"
	"strconv"
	"time"

	"github.com/google/uuid"
)

func generateDelivery(orderUID string) entity.Delivery {
	names := []string{
		"Иван Иванов",
		"Мария Смирнова",
		"Алексей Кузнецов",
		"Елена Попова",
		"Дмитрий Соколов",
		"Ольга Лебедева",
		"Сергей Новиков",
		"Анна Морозова",
		"Николай Павлов",
		"Татьяна Иванова",
	}

	phones := []string{
		"+7 910 000-0001",
		"+7 910 000-0002",
		"+7 910 000-0003",
		"+7 910 000-0004",
		"+7 910 000-0005",
		"+7 910 000-0006",
		"+7 910 000-0007",
		"+7 910 000-0008",
		"+7 910 000-0009",
		"+7 910 000-0010",
	}

	zips := []string{
		"101000",
		"102000",
		"103000",
		"104000",
		"105000",
		"106000",
		"107000",
		"108000",
		"109000",
		"110000",
	}

	cities := []string{
		"Москва",
		"Санкт-Петербург",
		"Новосибирск",
		"Екатеринбург",
		"Казань",
		"Нижний Новгород",
		"Челябинск",
		"Самара",
		"Омск",
		"Ростов-на-Дону",
	}

	addresses := []string{
		"ул. Ленина, д. 1, кв. 1",
		"ул. Пушкина, д. 5, кв. 10",
		"пр. Лермонтова, д. 12, кв. 3",
		"ул. Советская, д. 7, кв. 21",
		"наб. Речная, д. 2, кв. 4",
		"ул. Цветочная, д. 18, кв. 6",
		"пер. Гаражный, д. 9, кв. 2",
		"ул. Мира, д. 33, кв. 8",
		"ул. Островского, д. 4, кв. 12",
		"пр. Кирова, д. 20, кв. 14",
	}

	regions := []string{
		"Московская область",
		"Ленинградская область",
		"Новосибирская область",
		"Свердловская область",
		"Республика Татарстан",
		"Нижегородская область",
		"Челябинская область",
		"Самарская область",
		"Омская область",
		"Ростовская область",
	}

	emails := []string{
		"a@example.com",
		"b@example.com",
		"c@example.com",
		"d@example.com",
		"e@example.com",
		"f@example.com",
		"g@example.com",
		"h@example.com",
		"i@example.com",
		"j@example.com",
	}

	return entity.Delivery{
		OrderUID: orderUID,
		Name:     names[rand.Intn(10)],
		Phone:    phones[rand.Intn(10)],
		Zip:      zips[rand.Intn(10)],
		City:     cities[rand.Intn(10)],
		Address:  addresses[rand.Intn(10)],
		Region:   regions[rand.Intn(10)],
		Email:    emails[rand.Intn(10)],
	}

}

func generatePayment(OrderUID string) entity.Payment {
	return entity.Payment{
		OrderUID:     OrderUID,
		Transaction:  fmt.Sprintf("transaction-%d", time.Now().UnixNano()),
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       float64(rand.Intn(10000) + 200),
		PaymentDT:    int64(rand.Intn(90000000) + 10000000),
		Bank:         "alpha",
		DeliveryCost: float64(rand.Intn(100) + 100),
		GoodsTotal:   int64(rand.Intn(10000)),
		CustomFee:    0,
	}

}

func generateItems(trackNumber string) []entity.OrderItems {
	var orderItems []entity.OrderItems
	itemCount := rand.Intn(5) + 1
	for i := 0; i < itemCount; i++ {
		item := entity.Item{
			ItemID: int64(rand.Intn(500) + 1),
		}
		val := entity.OrderItems{
			Item:        item,
			RID:         fmt.Sprintf("rid-%d", time.Now().UnixNano()),
			TrackNumber: trackNumber,
			Sale:        int32(rand.Intn(30)),
			Size:        fmt.Sprintf("%d", rand.Intn(5)),
			TotalPrice:  int32(rand.Intn(10000) + 100), // по факту item.Price - sale, но тут просто сгенерируем тестовые данные
			Status:      202,
		}
		orderItems = append(orderItems, val)
	}
	return orderItems
}

func generateOrder() entity.Order {
	locales := []string{"en", "ru"}
	deliveryServices := []string{"test_delivery_service1", "test_delivery_service2", "test_delivery_service3", "test_delivery_service4", "test_delivery_service5"}
	id := uuid.New().String()
	orderUID := fmt.Sprintf("order-%s", id)
	trackNumber := fmt.Sprintf("WB-%s", id)
	customerID := fmt.Sprintf("customer-%s", id)
	order := entity.Order{
		OrderUID:          orderUID,
		TrackNumber:       trackNumber,
		Entry:             "WBIL",
		Delivery:          generateDelivery(orderUID),
		Payment:           generatePayment(orderUID),
		Items:             generateItems(trackNumber),
		Locale:            locales[rand.Intn(2)],
		InternalSignature: "",
		CustomerID:        customerID,
		DeliveryService:   deliveryServices[rand.Intn(5)],
		ShardKey:          int32(rand.Intn(10)),
		SmID:              int32(rand.Intn(100)),
		DateCreated:       time.Now(),
		OofShard:          strconv.Itoa(rand.Intn(10)),
	}
	return order
}

func generateOrdersChan(ctx context.Context) chan entity.Order {
	ordersChan := make(chan entity.Order)
	go func() {
		defer close(ordersChan)
		for {
			order := generateOrder()
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
