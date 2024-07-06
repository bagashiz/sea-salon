package repository

import (
	"context"

	"github.com/bagashiz/sea-salon/internal/postgres"
)

// DB is an interface that wraps the Querier interface and the ExecTX method for executing transactions.
type DB interface {
	postgres.Querier
	ExecTX(ctx context.Context, fn func(postgres.Querier) error) error
}
