# Используем официальный образ Go
FROM golang:alpine AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum
COPY go.mod go.sum ./
# Загружаем зависимости
RUN go mod download

# Копируем весь код приложения
COPY . .

RUN go build -o ./order_service ./cmd/app/main.go

FROM alpine:latest as runner

WORKDIR /app

COPY --from=builder /app/order_service ./order_service
COPY config/config.yml ./config/config.yml
COPY internal/database/repo/migrations/ ./internal/database/repo/sql/*.sql
COPY static ./static

# Открываем порт
EXPOSE 8081

# Запускаем приложение
CMD ["./order_service"]
