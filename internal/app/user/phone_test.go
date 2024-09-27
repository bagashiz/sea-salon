package user_test

import (
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

// TestNewPhoneNumber tests the NewPhoneNumber function.
func TestNewPhoneNumber(t *testing.T) {
	testCases := []struct {
		err   error
		desc  string
		input string
		want  user.PhoneNumber
	}{
		{desc: "valid_phone_number_0", input: "08123456789", want: user.PhoneNumber("08123456789"), err: nil},
		{desc: "valid_phone_number_1", input: "628123456789", want: user.PhoneNumber("628123456789"), err: nil},
		{desc: "empty_phone_number", input: "", want: "", err: user.ErrPhoneEmpty},
		{desc: "invalid_phone_number_0", input: "+628123456789", want: "", err: user.ErrPhoneInvalid},
		{desc: "invalid_phone_number_1", input: "0812-345-6789", want: "", err: user.ErrPhoneInvalid},
		{desc: "invalid_phone_number_2", input: "+1 (123) 456-7890", want: "", err: user.ErrPhoneInvalid},
		{desc: "invalid_phone_number_3", input: "0812345678a", want: "", err: user.ErrPhoneInvalid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := user.NewPhoneNumber(tc.input)
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
