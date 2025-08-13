package repo

import (
	"embed"
	"fmt"

	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func (pgRepo *PgRepo) ApplyMigrations() error {
	// Получаем *sql.DB из pgxpool.Pool
	db, err := pgRepo.postgres.GetSQLDB()
	if err != nil {
		return fmt.Errorf("ApplyMigrations pgRepo.postgres.GetSQLDB: %w", err)
	}
	defer db.Close()
	// Настраиваем goose
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("ApplyMigrations goose.SetDialect: %w", err)
	}
	// Применяем миграции
	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("ApplyMigrations goose.Up: %w", err)
	}

	return nil

}
