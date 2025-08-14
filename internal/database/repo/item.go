package repo

import (
	"context"
	"errors"
	"fmt"
	"orderService/internal/entity"

	"github.com/Masterminds/squirrel"
)

func (r *PgRepo) GetItemByItemID(ctx context.Context, itemID int64) (*entity.Item, error) {
	sql, args, err := r.postgres.Builder.Select("item_id", "chrt_id", "price", "name", "nm_id", "brand").
		From("items").Where(squirrel.Eq{"item_id": itemID}).ToSql()
	if err != nil {
		return nil, fmt.Errorf("GetShortOrderByOrderUID r.postgres.Builder.Select: %w", err)
	}
	var item entity.Item
	row := r.postgres.Pool.QueryRow(ctx, sql, args...)
	err = row.Scan(&item.ItemID, &item.ChrtID, &item.Price, &item.Name, &item.NmID, &item.Brand)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return nil, err
		}
		return nil, fmt.Errorf("GetShortOrderByOrderUID row.Scan: %w", err)
	}
	return &item, nil
}
