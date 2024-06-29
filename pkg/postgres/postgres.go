package postgres

import (
	"context"
	"fmt"

	"github.com/bagashiz/sea-salon/pkg/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// Storer extends the Querier interface with a Migrate method.
type Storer interface {
	Querier
	Migrate() error
}

// Store wraps the auto-generated Queries struct with a pgxpool.Pool.
type Store struct {
	*Queries
	pool *pgxpool.Pool
}

// NewStore creates a new DB instance using the provided config.
func NewStore(ctx context.Context, cfg *config.DB) (Storer, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return &Store{
		Queries: New(pool),
		pool:    pool,
	}, nil
}

// Migrate runs the goose migration tool to apply new migrations.
func (db *Store) Migrate() error {
	goose.SetBaseFS(migrationFS)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	conn := stdlib.OpenDBFromPool(db.pool)

	if err := goose.Up(conn, "migrations"); err != nil {
		return err
	}

	if err := conn.Close(); err != nil {
		return err
	}

	return nil
}
