package main

import (
	"log/slog"
	"orderService/config"
	_ "orderService/docs"
	"orderService/internal/app"
)

// @title Order Service
// @version 1.0
// @description API for Order Service
// @termsOfService http://swagger.io/terms/

// @host localhost:8081
// @BasePath /
func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		slog.Error("main config.New", slog.Any("error", err))
		return
	}
	app.Run(cfg)

}
