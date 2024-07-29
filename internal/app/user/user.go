package user

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user for the application.
type User struct {
	ID          uuid.UUID
	FullName    string
	Email       string
	Password    string
	PhoneNumber string
	Role        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewUser creates a new user.
func NewUser(fullName string, phoneNumber PhoneNumber, email Email, password Password, userRole UserRole) (*User, error) {
	return &User{
		FullName:    fullName,
		Email:       email.String(),
		Password:    password.String(),
		PhoneNumber: phoneNumber.String(),
		Role:        string(userRole),
	}, nil
}

// Create assigns a new UUID and CreateAt time to the user.
func (u *User) Create() {
	u.ID = uuid.New()
	u.CreatedAt = time.Now().Round(time.Second)
	u.UpdatedAt = time.Now().Round(time.Second)
}
