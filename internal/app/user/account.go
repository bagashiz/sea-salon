package user

import (
	"time"

	"github.com/google/uuid"
)

// Account represents a user for the application.
type Account struct {
	ID          uuid.UUID
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewAccount creates a new user account.
func NewAccount(
	fullName FullName,
	phoneNumber PhoneNumber,
	email Email,
	password Password,
	accountRole AccountRole,
) (*Account, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &Account{
		ID:          id,
		FullName:    fullName.String(),
		Email:       email.String(),
		Password:    password.String(),
		PhoneNumber: phoneNumber.String(),
		Role:        accountRole.String(),
		CreatedAt:   time.Now().Round(time.Second),
		UpdatedAt:   time.Now().Round(time.Second),
	}, nil
}
