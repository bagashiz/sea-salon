// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	DeleteUser(ctx context.Context, id pgtype.UUID) error
	GetAllUsers(ctx context.Context, arg GetAllUsersParams) ([]User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserById(ctx context.Context, id pgtype.UUID) (User, error)
	InsertUser(ctx context.Context, arg InsertUserParams) (User, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
}

var _ Querier = (*Queries)(nil)
