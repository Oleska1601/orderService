package repo

import (
	"context"
	"fmt"
	"orderService/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (r *PgRepo) SaveDelivery(ctx context.Context, tx pgx.Tx, delivery entity.Delivery) error {
	sql, args, err := r.postgres.Builder.Insert("deliveries").
		Columns("order_uid", "name", "phone", "zip", "city", "address", "region", "email").
		Values(delivery.OrderUID, delivery.Name, delivery.Phone, delivery.Zip, delivery.City, delivery.Address, delivery.Region, delivery.Email).ToSql()
	if err != nil {
		return fmt.Errorf("SaveDelivery r.postgres.Builder.Insert: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("SaveDelivery tx.Exec: %w", err)
	}
	return nil
}
