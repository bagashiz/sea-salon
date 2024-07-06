package user

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// emailMaxLength is the maximum length of an email address.
const emailMaxLength = 100

// validEmailSeq is a regular expression that matches a valid email address.
var validEmailSeq = regexp.MustCompile(`^[a-zA-Z0-9+._~\-]+@[a-zA-Z0-9+._~\-]+(\.[a-zA-Z0-9+._~\-]+)+$`)

// Email represents an email address.
type Email string

// NewEmail valiates and creates a new email address.
func NewEmail(email string) (Email, error) {
	if strings.TrimSpace(email) == "" {
		return "", ErrEmailEmpty
	}

	if rc := utf8.RuneCountInString(email); rc > emailMaxLength {
		return "", ErrEmailExceedsMaxLength
	}

	if !validEmailSeq.MatchString(email) {
		return "", ErrEmailInvalid
	}

	return Email(email), nil
}

// String returns the string representation of the email address.
func (e Email) String() string {
	return string(e)
}
