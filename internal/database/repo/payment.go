package repo

import (
	"context"
	"fmt"
	"orderService/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (r *PgRepo) SavePayment(ctx context.Context, tx pgx.Tx, payment entity.Payment) error {
	sql, args, err := r.postgres.Builder.Insert("payments").
		Columns("order_uid", "transaction", "request_id", "currency", "provider", "amount", "payment_dt", "bank", "delivery_cost", "goods_total", "custom_fee").
		Values(payment.OrderUID, payment.Transaction, payment.RequestID, payment.Currency, payment.Provider, payment.Amount, payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee).ToSql()
	if err != nil {
		return fmt.Errorf("SavePayment r.postgres.Builder.Insert: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("SavePayment tx.Exec: %w", err)
	}
	return nil
}
