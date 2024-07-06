package user

import (
	"strings"
)

// UserRole is an enum for user roles.
type UserRole string

const (
	Admin    UserRole = "admin"
	Customer UserRole = "customer"
)

// NewRole validates and creates a new user role.
func NewRole(role string) (UserRole, error) {
	switch strings.ToLower(role) {
	case string(Admin):
		return Admin, nil
	case string(Customer):
		return Customer, nil
	default:
		return "", ErrRoleInvalid
	}
}

// String returns the string representation of the user role.
func (r UserRole) String() string {
	return string(r)
}
