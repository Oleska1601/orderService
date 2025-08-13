package repo

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"orderService/internal/entity"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
)

func (r *PgRepo) GetOrderByOrderUID(ctx context.Context, orderUID string) (*entity.Order, error) {
	sqlString, args, err := r.postgres.Builder.Select("o.order_uid", "o.track_number", "o.entry", "o.locale", "o.internal_signature",
		"o.customer_id", "o.delivery_service", "o.shardkey", "o.sm_id", "o.date_created", "o.oof_shard",
		"d.name", "d.phone", "d.zip", "d.city", "d.address", "d.region", "d.email",
		"p.transaction", "p.request_id", "p.currency", "p.provider", "p.amount", "p.payment_dt", "p.bank", "p.delivery_cost", "p.goods_total", "p.custom_fee",
		// Подзапрос для items как JSON
		"(SELECT json_agg(json_build_object("+
			"'chrt_id', i.chrt_id,"+
			"'price', i.price,"+
			"'name', i.name,"+
			"'nm_id', i.nm_id,"+
			"'brand', i.brand,"+
			"'rid', oi.rid,"+
			"'sale', oi.sale,"+
			"'size', oi.size,"+
			"'total_price', oi.total_price,"+
			"'status', oi.status"+
			")) FROM items i JOIN order_items oi ON i.item_id = oi.item_id WHERE oi.track_number = o.track_number) AS items").
		From("orders as o").
		LeftJoin("deliveries as d on d.order_uid=o.order_uid").
		LeftJoin("payments as p on p.order_uid=o.order_uid").
		Where(squirrel.Eq{"o.order_uid": orderUID}).ToSql()

	if err != nil {
		return nil, fmt.Errorf("GetOrderByOrderUID r.postgres.Builder.Select: %w", err)
	}

	var order entity.Order
	var itemsJSON []byte
	row := r.postgres.Pool.QueryRow(ctx, sqlString, args...)
	err = row.Scan(
		&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale,
		&order.InternalSignature, &order.CustomerID, &order.DeliveryService,
		&order.ShardKey, &order.SmID, &order.DateCreated, &order.OofShard,

		// Delivery
		&order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip,
		&order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region,
		&order.Delivery.Email,

		// Payment
		&order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency,
		&order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDT,
		&order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal,
		&order.Payment.CustomFee,

		// Items (как JSON)
		&itemsJSON,
	)

	// Обрабатываем случай, когда запись не найдена
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("GetOrderByOrderUID row.Scan: %w", err)
	}
	if err := json.Unmarshal(itemsJSON, &order.Items); err != nil {
		return nil, fmt.Errorf("GetOrderByOrderUID json.Unmarshal: %w", err)
	}
	return &order, nil

}

func (r *PgRepo) GetOrdersWithCapacity(ctx context.Context, capacity int) ([]entity.Order, error) {
	sql, args, err := r.postgres.Builder.Select("o.order_uid", "o.track_number", "o.entry", "o.locale", "o.internal_signature",
		"o.customer_id", "o.delivery_service", "o.shardkey", "o.sm_id", "o.date_created", "o.oof_shard",
		"d.name", "d.phone", "d.zip", "d.city", "d.address", "d.region", "d.email",
		"p.transaction", "p.request_id", "p.currency", "p.provider", "p.amount", "p.payment_dt", "p.bank", "p.delivery_cost", "p.goods_total", "p.custom_fee",
		// Подзапрос для items как JSON
		"(SELECT json_agg(json_build_object("+
			"'chrt_id', i.chrt_id,"+
			"'price', i.price,"+
			"'name', i.name,"+
			"'nm_id', i.nm_id,"+
			"'brand', i.brand,"+
			"'rid', oi.rid,"+
			"'sale', oi.sale,"+
			"'size', oi.size,"+
			"'total_price', oi.total_price,"+
			"'status', oi.status"+
			")) FROM items i JOIN order_items oi ON i.item_id = oi.item_id WHERE oi.track_number = o.track_number) AS items").
		From("orders as o").
		LeftJoin("deliveries as d on d.order_uid=o.order_uid").
		LeftJoin("payments as p on p.order_uid=o.order_uid").OrderBy("o.date_created desc").
		Limit(uint64(capacity)).ToSql()

	if err != nil {
		return nil, fmt.Errorf("GetOrdersByQuantity r.postgres.Builder.Select: %w", err)
	}

	var orders []entity.Order
	rows, err := r.postgres.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("GetOrdersByQuantity r.postgres.Pool.Query: %w", err)
	}
	for rows.Next() {
		var order entity.Order
		var itemsJSON []byte
		err = rows.Scan(
			&order.OrderUID, &order.TrackNumber, &order.Entry, &order.Locale,
			&order.InternalSignature, &order.CustomerID, &order.DeliveryService,
			&order.ShardKey, &order.SmID, &order.DateCreated, &order.OofShard,

			// Delivery
			&order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip,
			&order.Delivery.City, &order.Delivery.Address, &order.Delivery.Region,
			&order.Delivery.Email,

			// Payment
			&order.Payment.Transaction, &order.Payment.RequestID, &order.Payment.Currency,
			&order.Payment.Provider, &order.Payment.Amount, &order.Payment.PaymentDT,
			&order.Payment.Bank, &order.Payment.DeliveryCost, &order.Payment.GoodsTotal,
			&order.Payment.CustomFee,

			// Items (как JSON)
			&itemsJSON,
		)
		if err != nil {
			return nil, fmt.Errorf("GetOrdersByQuantity row.Scan: %w", err)
		}
		if err := json.Unmarshal(itemsJSON, &order.Items); err != nil {
			return nil, fmt.Errorf("GetOrdersByQuantity: json.Unmarshal %w", err)
		}
		orders = append(orders, order)
	}

	return orders, nil

}

func (r *PgRepo) SaveOrder(ctx context.Context, tx pgx.Tx, order entity.Order) error {
	sql, args, err := r.postgres.Builder.Insert("orders").
		Columns("order_uid", "track_number", "entry", "locale", "internal_signature", "customer_id", "delivery_service", "shardkey", "sm_id", "date_created", "oof_shard").
		Values(order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature, order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated, order.OofShard).ToSql()
	if err != nil {
		return fmt.Errorf("SavePayment r.postgres.Builder.Insert: %w", err)
	}

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("SavePayment tx.Exec: %w", err)
	}
	return nil
}

func (r *PgRepo) CheckOrderExistsByOrderUID(ctx context.Context, orderUID string) (bool, error) {
	var isExists bool
	sql, args, err := r.postgres.Builder.Select("1").Prefix("select exists (").
		From("orders").Where(squirrel.Eq{"order_uid": orderUID}).Suffix(")").ToSql()
	if err != nil {
		return false, fmt.Errorf("GetShortOrderByOrderUID r.postgres.Builder.Select: %w", err)
	}
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&isExists)
	if err != nil {
		return false, fmt.Errorf("GetShortOrderByOrderUID row.Scan: %w", err)
	}
	return isExists, nil
}
