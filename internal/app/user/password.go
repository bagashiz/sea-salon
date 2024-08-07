package user

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"golang.org/x/crypto/bcrypt"
)

type Password string

const (
	// passwordMinLength is the minimum length of a password.
	passwordMinLength = 8
	// passwordMaxLength is the maximum length of a password.
	passwordMaxLength = 64
)

// passwordRegex is a regular expression for alphanumeric characters and symbols.
var passwordRegex = regexp.MustCompile("^[a-zA-Z0-9!@#$%^&*()\\-_=+\\[\\]{}|;:'\",.<>?/~`]+$")

// NewPassword validates and creates a new password.
func NewPassword(password string) (Password, error) {
	if strings.TrimSpace(password) == "" {
		return "", ErrPasswordEmpty
	}

	rc := utf8.RuneCountInString(password)

	if rc < passwordMinLength {
		return "", ErrPasswordTooShort
	}

	if rc > passwordMaxLength {
		return "", ErrPasswordTooLong
	}

	if !passwordRegex.MatchString(password) {
		return "", ErrPasswordInvalid
	}

	return Password(password), nil
}

// Hash returns the hashed password with bcrypt typed as Password.
func (p Password) Hash() (Password, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return Password(string(hash)), nil
}

// Compare returns true if the password matches the hash.
func (p Password) Compare(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(p))
	return err == nil
}

// String returns the string representation of the password.
func (p Password) String() string {
	return string(p)
}
