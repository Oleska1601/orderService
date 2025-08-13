package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
)

const (
	_defaultMaxPoolSize    = 1
	_defaultMaxConnAttemts = 10
	_defaultMaxConnTimeout = time.Second
)

type Postgres struct {
	maxPoolSize     int
	maxConnAttempts int
	maxConnTimeout  time.Duration
	Builder         squirrel.StatementBuilderType
	Pool            *pgxpool.Pool
}

func NewPostgres(pgUrl string, opts ...Option) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:     _defaultMaxPoolSize,
		maxConnAttempts: _defaultMaxConnAttemts,
		maxConnTimeout:  _defaultMaxConnTimeout,
	}
	for _, opt := range opts {
		opt(pg)
	}
	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	pgxConfig, err := pgxpool.ParseConfig(pgUrl)
	if err != nil {
		return nil, err
	}
	pgxConfig.MaxConns = int32(pg.maxPoolSize)
	for pg.maxConnAttempts > 0 {
		slog.Debug("NewPostgres postgres is trying to connect", slog.Int("attempts left", pg.maxConnAttempts))
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), pgxConfig)
		if err == nil {
			break
		}
		time.Sleep(pg.maxConnTimeout)
		pg.maxConnAttempts--
	}
	if err != nil {
		return nil, fmt.Errorf("NewPostgres connAttempts == 0: %w", err)
	}

	return pg, nil
}

func (pg *Postgres) Close() {
	if pg.Pool != nil {
		pg.Pool.Close()
	}

}

func (pg *Postgres) GetSQLDB() (*sql.DB, error) {
	conn, err := pg.Pool.Acquire(context.Background())
	if err != nil {
		return nil, fmt.Errorf("GetSQLDB pg.Pool.Acquire: %w", err)
	}
	defer conn.Release()

	connConfig := conn.Conn().Config()
	db := stdlib.OpenDB(*connConfig)
	return db, nil
}
