package user

import (
	"regexp"
	"strings"
)

type PhoneNumber string

var (
	// validPhoneSeq is a regular expression that matches a valid phone number.
	validPhoneSeq = regexp.MustCompile(`^[0-9]+$`)
)

// NewPhone validates and creates a new phone number.
func NewPhone(number string) (PhoneNumber, error) {
	if strings.TrimSpace(number) == "" {
		return "", ErrPhoneEmpty
	}

	if !validPhoneSeq.MatchString(number) {
		return "", ErrPhoneInvalid
	}

	return PhoneNumber(number), nil
}

// String returns the string representation of the phone number.
func (p PhoneNumber) String() string {
	return string(p)
}
