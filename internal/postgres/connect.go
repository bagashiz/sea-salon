package postgres

import (
	"context"
	"fmt"

	"github.com/bagashiz/sea-salon/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

// DB wraps the Querier interface and a pgxpool.Pool instance.
type DB struct {
	Querier
	*pgxpool.Pool
}

// NewDB creates a new DB instance using the provided config.
func NewDB(ctx context.Context, cfg *config.DB) (*DB, error) {
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

	return &DB{Querier: New(pool), Pool: pool}, nil
}

// Migrate runs the goose migration tool to apply new migrations.
func (d *DB) Migrate(dialect string) error {
	goose.SetBaseFS(migrationFS)

	if err := goose.SetDialect(dialect); err != nil {
		return err
	}

	db := stdlib.OpenDBFromPool(d.Pool)
	defer db.Close()

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}

// ExecTX wraps the provided function in a transaction and executes it.
func (d *DB) ExecTX(ctx context.Context, fn func(Querier) error) error {
	tx, err := d.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	if err := fn(New(tx)); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return rbErr
		}

		return err
	}

	return tx.Commit(ctx)
}
