package repo

import (
	"context"
	"fmt"
	"orderService/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (r *PgRepo) BeginTx(ctx context.Context) (pgx.Tx, error) {
	tx, err := r.postgres.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("BeginTx r.postgres.Pool.Begin: %w", err)
	}
	return tx, nil
}

func (r *PgRepo) SaveMessage(ctx context.Context, tx pgx.Tx, order entity.Order) error {

	if err := r.SaveOrder(ctx, tx, order); err != nil {
		return fmt.Errorf("SaveMessage r.SaveOrder: %w", err)
	}
	if err := r.SaveDelivery(ctx, tx, order.Delivery); err != nil {
		return fmt.Errorf("SaveMessage r.SaveDelivery: %w", err)
	}

	if err := r.SavePayment(ctx, tx, order.Payment); err != nil {
		return fmt.Errorf("SaveMessage r.SavePayment: %w", err)
	}

	if err := r.SaveOrderItems(ctx, tx, order.Items); err != nil {
		return fmt.Errorf("SaveMessage r.SaveOrderItems: %w", err)
	}

	return nil
}
