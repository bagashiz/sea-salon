package user

import (
	"context"

	"github.com/google/uuid"
)

// Reader is the interface that provides methods to read user data from the storage.
type Reader interface {
	GetAccountByID(ctx context.Context, id uuid.UUID) (*Account, error)
	GetAccountByEmail(ctx context.Context, email string) (*Account, error)
	ListAccounts(ctx context.Context, limit, offset int) ([]*Account, error)
}

// Writer is the interface that provides methods to write user data to the storage.
type Writer interface {
	AddAccount(ctx context.Context, u *Account) error
	UpdateAccount(ctx context.Context, u *Account) error
	DeleteAccount(ctx context.Context, id uuid.UUID) error
}

// ReadWriter is the interface that combines Reader and Writer interfaces for user data.
type ReadWriter interface {
	Reader
	Writer
}
