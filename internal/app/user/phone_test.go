package user_test

import (
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

func TestNewPhone(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  user.PhoneNumber
		err   error
	}{
		{desc: "valid phone number", input: "08123456789", want: user.PhoneNumber("08123456789"), err: nil},
		{desc: "valid phone number", input: "628123456789", want: user.PhoneNumber("628123456789"), err: nil},
		{desc: "empty phone number", input: "", want: "", err: user.ErrPhoneEmpty},
		{desc: "invalid phone number", input: "+628123456789", want: "", err: user.ErrPhoneInvalid},
		{desc: "invalid phone number", input: "0812-345-6789", want: "", err: user.ErrPhoneInvalid},
		{desc: "invalid phone number", input: "+1 (123) 456-7890", want: "", err: user.ErrPhoneInvalid},
		{desc: "invalid phone number", input: "0812345678a", want: "", err: user.ErrPhoneInvalid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := user.NewPhone(tc.input)
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
