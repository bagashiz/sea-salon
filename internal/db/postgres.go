package db

import (
	"context"
	"fmt"

	"github.com/bagashiz/sea-salon/pkg/config"
	"github.com/bagashiz/sea-salon/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// Postgres represents a Postgres database connection.
type Postgres struct {
	pool    *pgxpool.Pool
	Queries *postgres.Queries
}

// NewPostgres creates a new Postgres instance.
func NewPostgres(ctx context.Context, cfg *config.DB) (*Postgres, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &Postgres{
		pool:    pool,
		Queries: postgres.New(pool),
	}, nil
}

// Migrate runs the database migrations.
func (p *Postgres) Migrate() error {
	goose.SetBaseFS(postgres.MigrationFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(p.pool)

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	if err := db.Close(); err != nil {
		return err
	}

	return nil
}
