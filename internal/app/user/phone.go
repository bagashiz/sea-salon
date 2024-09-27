package user

import (
	"regexp"
	"strings"
)

type PhoneNumber string

// validPhoneSeq is a regular expression that matches a valid phone number.
var validPhoneSeq = regexp.MustCompile(`^[0-9]+$`)

// NewPhoneNumber validates and creates a new phone number.
func NewPhoneNumber(number string) (PhoneNumber, error) {
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
