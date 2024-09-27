package user

import "strings"

// AccountRole is an enum for user account roles.
type AccountRole string

const (
	Admin    AccountRole = "admin"
	Customer AccountRole = "customer"
)

// NewAccountRole validates and creates a new user account role.
func NewAccountRole(role string) (AccountRole, error) {
	switch strings.ToLower(role) {
	case string(Admin):
		return Admin, nil
	case string(Customer):
		return Customer, nil
	default:
		return "", ErrRoleInvalid
	}
}

// String returns the string representation of the user account role.
func (r AccountRole) String() string {
	return string(r)
}
