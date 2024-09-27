package user

import "strings"

// FullName represents a user's full name.
type FullName string

// NewFullName validates and creates a new full name.
func NewFullName(name string) (FullName, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrFullNameEmpty
	}
	return FullName(name), nil
}

// String returns the string representation of the full name.
func (n FullName) String() string {
	return string(n)
}
