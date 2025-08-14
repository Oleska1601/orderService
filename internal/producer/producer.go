package producer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"orderService/internal/database/repo"
	"orderService/pkg/logger"
	"time"

	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
	pgRepo repo.PgRepoInterface
	logger logger.LoggerInterface
}

func NewProducer(brokers []string, topic string, pgRepo repo.PgRepoInterface, logger logger.LoggerInterface) *Producer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      brokers,
		Topic:        topic,
		BatchSize:    10,
		BatchTimeout: 2 * time.Second,
		RequiredAcks: -1, //сообщение считается отправленным, когда его записали все реплики брокера
	})
	return &Producer{
		writer: w,
		pgRepo: pgRepo,
		logger: logger,
	}
}

func (p *Producer) RunProducer(ctx context.Context) {
	p.logger.Info("start producer")
	defer p.StopProducer()
	ordersChan := p.generateOrdersChan(ctx)
	for {
		select {
		case <-ctx.Done():
			p.logger.Info("RunProducer ctx.Done", slog.String("message", "stop producer due to context"))
			return
		case order := <-ordersChan:
			value, err := json.Marshal(order)
			if err != nil {
				p.logger.Error("RunProducer json.Marshal", slog.Any("error", err))
				continue
			}
			msg := kafka.Message{
				Key:   []byte(order.OrderUID),
				Value: value,
			}
			err = p.writer.WriteMessages(ctx, msg)
			if err != nil {
				if errors.Is(err, context.Canceled) {
					p.logger.Info("RunProducer ctx.Done", slog.String("message", "stop producer due to context"))
					return
				}
				p.logger.Error("RunProducer p.writer.WriteMessages", slog.Any("error", err))
				continue
			}
			p.logger.Info("RunProducer", slog.String("message", "produce order is successful"),
				slog.Any("offset", msg.Offset),
				slog.Int("partition", msg.Partition),
				slog.String("message key", string(msg.Key)))
		}
	}

}

func (p *Producer) StopProducer() error {
	if err := p.writer.Close(); err != nil {
		return fmt.Errorf("StopProducer p.writer.Close: %w", err)
	}
	p.logger.Info("StopProducer", slog.Any("message", "stop producer is successful"))
	return nil
}
