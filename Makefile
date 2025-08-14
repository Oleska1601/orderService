.PHONY: build run docker-build docker-run migrate swagger

build:
	go build -o bin/app ./cmd/app

run: build
	./bin/app

docker-build:
	docker build -t order_service .

docker-run: docker-build
	docker-compose up --build

migrate:
	go run cmd/app/main.go migrate

swagger:
	swag init -g cmd/app/main.go -o docs