package user_test

import (
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

// TestNewEmail tests the NewEmail function.
func TestNewEmail(t *testing.T) {
	testCases := []struct {
		err   error
		desc  string
		input string
		want  user.Email
	}{
		{desc: "valid_email_0", input: "test@example.com", want: user.Email("test@example.com"), err: nil},
		{desc: "valid_email_1", input: "test+go-salon@example.com", want: user.Email("test+go-salon@example.com"), err: nil},
		{desc: "no_input", input: "", want: user.Email(""), err: user.ErrEmailEmpty},
		{desc: "exceeds_max_length", input: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa@example.com", want: user.Email(""), err: user.ErrEmailExceedsMaxLength},
		{desc: "contains_whitespace_0", input: "test @example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains_whitespace_1", input: "test@example.com\t", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains_whitespace_3", input: "\rtest@example.com ", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains_quotes_0", input: "test\"@example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains_quotes_1", input: "\"test@example.com\"", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "contains_quotes_2", input: "'test@example.com'", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid_characters_0", input: "te#st@example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid_characters_1", input: "test!@example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid_characters_2", input: "te=st@example.com", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid_format_0", input: "test@", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid_format_1", input: "@example", want: user.Email(""), err: user.ErrEmailInvalid},
		{desc: "invalid_format_2", input: "test@example", want: user.Email(""), err: user.ErrEmailInvalid},
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
