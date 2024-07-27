package user

import (
	"errors"
	"fmt"
)

var (
	// list of errors for full name validation
	ErrFullNameEmpty = errors.New("name cannot be empty")
	// list of errors for email validation
	ErrEmailEmpty            = errors.New("email cannot be empty")
	ErrEmailExceedsMaxLength = fmt.Errorf("email cannot exceed %d characters", emailMaxLength)
	ErrEmailInvalid          = errors.New("email must be a valid email address")
	// list of errors for password validation
	ErrPasswordEmpty    = errors.New("password cannot be empty")
	ErrPasswordTooShort = fmt.Errorf("password cannot be less %d characters long", passwordMinLength)
	ErrPasswordTooLong  = fmt.Errorf("password cannot exceed %d characters", passwordMaxLength)
	ErrPasswordInvalid  = errors.New("password must contain only alphanumeric characters and symbols")
	ErrPasswordMismatch = errors.New("passwords do not match")
	// list of errors for phone number validation
	ErrPhoneEmpty   = errors.New("phone number cannot be empty")
	ErrPhoneInvalid = errors.New("phone number must contain only numbers")
	// list of errors for role validation
	ErrRoleInvalid = errors.New("user role must be either \"admin\" or \"customer\"")
	// list of errors for user validation
	ErrIDInvalid = errors.New("user ID must be a valid UUIDv4")
	// list of errors for user repository
	ErrUserInvalid  = errors.New("user data is invalid")
	ErrUserExists   = errors.New("user already exists")
	ErrUserNotFound = errors.New("user not found")
)
