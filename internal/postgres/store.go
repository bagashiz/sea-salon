package postgres

import (
	"context"
	"fmt"

	"github.com/bagashiz/sea-salon/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// Storer extends the Querier interface with a Migrate method.
type Storer interface {
	Querier
	Migrate(dialect string) error
}

// Store wraps the auto-generated Queries struct with a pgxpool.Pool.
type Store struct {
	*Queries
	pool *pgxpool.Pool
}

// NewStore creates a new DB instance using the provided config.
func NewStore(ctx context.Context, cfg *config.DB) (Storer, error) {
	dsn := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.Type, cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name,
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
func (s *Store) Migrate(dialect string) error {
	goose.SetBaseFS(migrationFS)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(s.pool)
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}
