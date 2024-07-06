package user

import "context"

// Reader is the interface that provides methods to read user data from the storage.
type Reader interface {
	GetUserByID(ctx context.Context, id string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	ListUsers(ctx context.Context, limit, offset int) ([]*User, error)
}

// Writer is the interface that provides methods to write user data to the storage.
type Writer interface {
	CreateUser(ctx context.Context, u *User) error
	UpdateUser(ctx context.Context, u *User) error
	DeleteUser(ctx context.Context, id string) error
}

// ReadWriter is the interface that combines Reader and Writer interfaces.
type ReadWriter interface {
	Reader
	Writer
}
