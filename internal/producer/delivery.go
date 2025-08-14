package producer

import (
	"math/rand"
	"orderService/internal/entity"
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
	name := names[rand.Intn(10)]
	phone := phones[rand.Intn(10)]
	zip := zips[rand.Intn(10)]
	city := cities[rand.Intn(10)]
	address := addresses[rand.Intn(10)]
	region := regions[rand.Intn(10)]
	email := emails[rand.Intn(10)]
	return entity.Delivery{
		OrderUID: orderUID,
		Name:     name,
		Phone:    phone,
		Zip:      zip,
		City:     city,
		Address:  address,
		Region:   region,
		Email:    email,
	}

}
