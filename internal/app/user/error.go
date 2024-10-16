package user

import "fmt"

// UserError is a custom error for user domain logic.
type UserError struct {
	message string
}

// Error returns the error message for the UserError type.
func (ue *UserError) Error() string {
	return ue.message
}

var (
	// list of errors for full name validation
	ErrFullNameEmpty = &UserError{"full name cannot be empty"}
	// list of errors for email validation
	ErrEmailEmpty            = &UserError{"email cannot be empty"}
	ErrEmailExceedsMaxLength = &UserError{fmt.Sprintf("email cannot exceed %d characters", emailMaxLength)}
	ErrEmailInvalid          = &UserError{"email must be a valid email address"}
	// list of errors for password validation
	ErrPasswordEmpty    = &UserError{"password cannot be empty"}
	ErrPasswordTooShort = &UserError{fmt.Sprintf("password cannot be less than %d characters long", passwordMinLength)}
	ErrPasswordTooLong  = &UserError{fmt.Sprintf("password cannot exceed %d characters", passwordMaxLength)}
	ErrPasswordInvalid  = &UserError{"password must contain only alphanumeric characters and symbols"}
	ErrPasswordMismatch = &UserError{"passwords do not match"}
	// list of errors for phone number validation
	ErrPhoneEmpty   = &UserError{"phone number cannot be empty"}
	ErrPhoneInvalid = &UserError{"phone number must contain only numbers"}
	// list of errors for role validation
	ErrRoleInvalid = &UserError{"account role must be either \"admin\" or \"customer\""}
	// list of errors for user repository
	ErrAccountInvalid  = &UserError{"account data is invalid"}
	ErrAccountExists   = &UserError{"account already exists"}
	ErrAccountNotFound = &UserError{"account not found"}
)
