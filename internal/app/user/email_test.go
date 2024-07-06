package user_test

import (
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

// TestNewEmail tests the NewEmail function.
func TestNewEmail(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  user.Email
		err   error
	}{
		{desc: "valid email", input: "test@example.com", want: user.Email("test@example.com"), err: nil},
		{desc: "valid email", input: "test+go-salon@example.com", want: user.Email("test+go-salon@example.com"), err: nil},
		{desc: "no input", input: "", want: user.Email(""), err: user.ErrEmailEmpty},
		{desc: "exceeds max length", input: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com", want: user.Email(""), err: user.ErrEmailExceedsMaxLength},
		{desc: "contains whitespace", input: "test @example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains whitespace", input: "test@example.com\t", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains whitespace", input: "\rtest@example.com ", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains quotes", input: "test\"@example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains quotes", input: "\"test@example.com\"", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains quotes", input: "'test@example.com'", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid characters", input: "te#st@example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid characters", input: "test!@example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid characters", input: "te=st@example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid format", input: "test@", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid format", input: "@example", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid format", input: "test@example", want: user.Email(""), err: user.ErrEmailInvalid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := user.NewEmail(tc.input)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}

			if got != tc.want {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.want, got)
				return
			}
		})
	}
}
