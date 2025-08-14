package app

import (
	"context"
	"log/slog"
	"orderService/config"
	"orderService/internal/cache"
	"orderService/internal/consumer"
	"orderService/internal/controller"
	"orderService/internal/database/repo"
	"orderService/internal/producer"
	"orderService/internal/usecase"
	"orderService/pkg/logger"
	"orderService/pkg/postgres"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(cfg *config.Config) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logger := logger.NewLogger(cfg.Logger.Level)
	pg, err := postgres.NewPostgres(cfg.DB.PgUrl, postgres.SetMaxPoolSize(cfg.DB.MaxPoolSize))
	if err != nil {
		logger.Error("main postgres.NewPostgres", slog.Any("error", err))
		return
	}
	defer pg.Close()
	pgRepo := repo.NewRepo(pg)
	if err := pgRepo.ApplyMigrations(); err != nil {
		logger.Error("main pgRepo.ApplyMigrations", slog.Any("error", err))

		return
	}
	logger.Info("apply migrations successful")

	cache := cache.New(cfg.Cache.Capacity)
	logger.Info("initialise cashe successful")
	usecase, err := usecase.NewUsecase(cache, pgRepo, logger)
	if err != nil {
		logger.Error("main usecase.NewUsecase", slog.Any("error", err))
		return
	}
	server := controller.New(usecase, logger)
	consumer := consumer.NewConsumer(cfg.Kafka.Brokers, cfg.Kafka.Topic, pgRepo, logger)
	producer := producer.NewProducer(cfg.Kafka.Brokers, cfg.Kafka.Topic, pgRepo, logger)
	go consumer.RunConsumer(ctx)
	go producer.RunProducer(ctx)
	go server.Run(cfg.Server.Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	slog.Info("shutting down server...")
	cancel()

	time.Sleep(1 * cfg.Server.ShutdownTimeout) // Даем время на завершение операций

}
