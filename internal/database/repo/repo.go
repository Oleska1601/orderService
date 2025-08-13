package repo

import (
	"context"
	"orderService/internal/entity"
	"orderService/pkg/postgres"

	"github.com/jackc/pgx/v5"
)

type PgRepoInterface interface {
	ApplyMigrations() error
	GetOrderByOrderUID(context.Context, string) (*entity.Order, error)
	GetOrdersWithCapacity(context.Context, int) ([]entity.Order, error)
	SaveDelivery(context.Context, pgx.Tx, entity.Delivery) error
	SaveMessage(context.Context, pgx.Tx, entity.Order) error
	SaveOrder(context.Context, pgx.Tx, entity.Order) error
	SaveOrderItems(context.Context, pgx.Tx, []entity.OrderItems) error
	SavePayment(context.Context, pgx.Tx, entity.Payment) error
	BeginTx(context.Context) (pgx.Tx, error)
	CheckOrderExistsByOrderUID(context.Context, string) (bool, error)
}

type PgRepo struct {
	postgres *postgres.Postgres
}

func NewRepo(postgres *postgres.Postgres) *PgRepo {
	return &PgRepo{
		postgres: postgres,
	}
}
