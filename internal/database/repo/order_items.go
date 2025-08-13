package repo

import (
	"context"
	"fmt"
	"orderService/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (r *PgRepo) SaveOrderItems(ctx context.Context, tx pgx.Tx, orderItems []entity.OrderItems) error {
	for _, oi := range orderItems {
		sql, args, err := r.postgres.Builder.Insert("order_items").
			Columns("rid", "item_id", "track_number", "sale", "size", "total_price", "status").
			Values(oi.RID, oi.Item.ItemID, oi.TrackNumber, oi.Sale, oi.Size, oi.TotalPrice, oi.Status).ToSql()
		if err != nil {
			return fmt.Errorf("SaveOrderItems r.postgres.Builder.Insert: %w", err)
		}
		_, err = tx.Exec(ctx, sql, args...)
		if err != nil {
			return fmt.Errorf("SaveOrderItems tx.Exec: %w", err)
		}
	}

	return nil
}
