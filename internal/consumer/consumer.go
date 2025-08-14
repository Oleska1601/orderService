package consumer

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"orderService/internal/cache"
	"orderService/internal/database/repo"
	"orderService/pkg/logger"

	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader *kafka.Reader
	cache  *cache.LRUCache
	pgRepo repo.PgRepoInterface
	logger logger.LoggerInterface
}

func NewConsumer(brokers []string, topic string, cache *cache.LRUCache, pgRepo repo.PgRepoInterface, logger logger.LoggerInterface) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		GroupID:     "order-service",
		StartOffset: kafka.FirstOffset,
	})
	c := &Consumer{
		reader: r,
		cache:  cache,
		pgRepo: pgRepo,
		logger: logger,
	}
	return c
}

func (c *Consumer) RunConsumer(ctx context.Context) {
	c.logger.Info("start consumer")
	defer c.StopConsumer()
	for {
		select {
		case <-ctx.Done():
			c.logger.Info("Consume ctx.Done", slog.String("message", "stop consumer due to context"))
			return
		default:
			msg, err := c.reader.ReadMessage(ctx)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					c.logger.Info("Consume ctx.Done", slog.String("message", "stop consumer due to context"))
					return
				}
				c.logger.Error("Consume r.ReadMessage", slog.Any("error", err))
				continue
			}
			c.ConsumerProcessMessage(ctx, msg)

		}
	}
}

func (c *Consumer) StopConsumer() error {
	if err := c.reader.Close(); err != nil {
		return fmt.Errorf("StopConsumer p.reader.Close: %w", err)
	}
	c.logger.Info("StopConsumer", slog.Any("message", "stop consumer is successful"))
	return nil
}
