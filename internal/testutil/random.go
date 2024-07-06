package testutil

import (
	"fmt"
	"math/rand/v2"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers  = "0123456789"
)

func init() {
	rand.New(rand.NewPCG(uint64(time.Now().UnixNano()), uint64(time.Now().UnixNano())))
}

// RandomString generates a random alphabetical string of length n.
func RandomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = alphabet[rand.IntN(len(alphabet))]
	}
	return string(b)
}

// RandomDigits generates a random numerical string of length n.
func RandomDigits(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numbers[rand.IntN(len(numbers))]
	}
	return string(b)
}

// RandomAlphaNumeric generates a random alphanumeric string of length n.
func RandomAlphaNumeric(n int) string {
	b := make([]byte, n)
	alphanumeric := alphabet + numbers
	for i := range b {
		b[i] = alphanumeric[rand.IntN(len(alphanumeric))]
	}
	return string(b)
}

// Random FullName generates a random full name.
func RandomFullName() string {
	return fmt.Sprintf("%s %s", RandomString(5), RandomString(8))
}

// RandomEmail generates a random email address.
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomAlphaNumeric(8))
}

// RandomPhoneNumber generates a random phone number.
func RandomPhoneNumber() string {
	return fmt.Sprintf("08%s", RandomDigits(9))
}
