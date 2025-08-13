package consumer

import (
	"context"
	"encoding/json"
	"log/slog"
	"orderService/internal/entity"

	"github.com/segmentio/kafka-go"
)

func (c *Consumer) ConsumerProcessMessage(ctx context.Context, msg kafka.Message) {
	var order entity.Order

	// 1. Парсинг JSON
	if err := json.Unmarshal(msg.Value, &order); err != nil {
		c.logger.Error("ConsumerProcessMessage json.Unmarshal",
			slog.Any("error", err),
			slog.String("message key", string(msg.Key)))
		return
	}
	c.logger.Debug("ConsumerProcessMessage json.Unmarshal", slog.String("message", "marshal order"), slog.String("message key", string(msg.Key)))

	// 2. Валидация
	if err := order.Validate(); err != nil {
		c.logger.Error("ConsumerProcessMessage order.Validate", slog.Any("error", err),
			slog.String("message key", string(msg.Key)))
		return
	}
	c.logger.Debug("ConsumerProcessMessage order.Validate", slog.String("message", "order is valid"), slog.String("message key", string(msg.Key)))

	//3. Проверка дубликата
	exists, err := c.pgRepo.CheckOrderExistsByOrderUID(ctx, order.OrderUID)
	if err != nil {
		c.logger.Error("ConsumerProcessMessage c.pgRepo.CheckOrderExistsByOrderUID", slog.Any("error", err),
			slog.String("message key", string(msg.Key)))
		return
	}
	if exists { //если есть -> фиксирует оффсет
		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			c.logger.Error("ConsumerProcessMessage c.reader.CommitMessages", slog.Any("error", err),
				slog.String("message key", string(msg.Key)))

		}
		return
	}

	c.logger.Debug("ConsumerProcessMessage c.pgRepo.CheckOrderExistsByOrderUID", slog.String("message", "no dublicates found"), slog.String("message key", string(msg.Key)))

	//4. Начало транзакции
	tx, err := c.pgRepo.BeginTx(ctx)
	if err != nil {
		c.logger.Error("ConsumerProcessMessage c.pgRepo.BeginTx", slog.Any("error", err),
			slog.String("message key", string(msg.Key)))
		return
	}
	c.logger.Debug("ConsumerProcessMessage c.pgRepo.BeginTx", slog.String("message", "begin transaction"), slog.String("message key", string(msg.Key)))

	defer func() {
		if err != nil {
			if rollBackErr := tx.Rollback(ctx); rollBackErr != nil {
				c.logger.Error("ConsumerProcessMessage tx.Rollback", slog.Any("error", err), slog.String("message key", string(msg.Key)))
			} // гарантированный откат при падении / ошибки
		}
		c.logger.Debug("ConsumerProcessMessage tx.Rollback", slog.String("message", "rollback transaction"), slog.String("message key", string(msg.Key)))
	}()

	//5. Сохранение в БД
	if err := c.pgRepo.SaveMessage(ctx, tx, order); err != nil {
		c.logger.Error("ConsumerProcessMessage u.pgRepo.SaveMessage", slog.Any("error", err),
			slog.String("message key", string(msg.Key)))
		return
	}
	c.logger.Debug("ConsumerProcessMessage c.pgRepo.SaveMessage", slog.String("message", "save message in db"), slog.String("message key", string(msg.Key)))

	//6. Завершение транзакции
	if err := tx.Commit(ctx); err != nil {
		c.logger.Error("ConsumerProcessMessage tx.Commit", slog.Any("error", err),
			slog.String("message key", string(msg.Key)))
	}

	c.logger.Debug("ConsumerProcessMessage tx.Commit", slog.String("message", "commit transaction"), slog.String("message key", string(msg.Key)))

	if err := c.reader.CommitMessages(ctx, msg); err != nil {
		c.logger.Error("ConsumerProcessMessage c.reader.CommitMessages", slog.Any("error", err),
			slog.String("message key", string(msg.Key)))
		return
	}
	c.logger.Debug("ConsumerProcessMessage c.reader.CommitMessages", slog.String("message", "commit message"), slog.String("message key", string(msg.Key)))

	c.logger.Info("ConsumerProcessMessage", slog.String("message", "process message is successful"),
		slog.Any("offset", msg.Offset),
		slog.Int("partition", msg.Partition),
		slog.String("message key", string(msg.Key)))

}
