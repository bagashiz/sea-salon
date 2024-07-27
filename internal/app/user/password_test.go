package user_test

import (
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

// TestPassword tests the NewPassword and Compare methods of the Password type.
func TestPassword(t *testing.T) {
	testCases := []struct {
		err        error
		desc       string
		input      string
		wrongInput string
		isMatch    bool
	}{
		{desc: "valid password", input: "test1234", wrongInput: "", isMatch: true, err: nil},
		{desc: "valid password", input: "#tes+t1@-4e2=34$", wrongInput: "", isMatch: true, err: nil},
		{desc: "valid password", input: "$2a$10$4tS9MQtS6l/9PWY.MiR8O.3.yFKHvaB34kpQVGgSVnGrla6ztOaam", wrongInput: "", isMatch: true, err: nil},
		{desc: "empty password", input: "", wrongInput: "", isMatch: false, err: user.ErrPasswordEmpty},
		{desc: "password too short", input: "test", wrongInput: "", isMatch: false, err: user.ErrPasswordTooShort},
		{desc: "password too long", input: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", wrongInput: "", isMatch: false, err: user.ErrPasswordTooLong},
		{desc: "invalid password", input: "test\t1234", wrongInput: "", isMatch: false, err: user.ErrPasswordInvalid},
		{desc: "invalid password", input: " test1234 ", wrongInput: "", isMatch: false, err: user.ErrPasswordInvalid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			password, err := user.NewPassword(tc.input)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}

			testPassword := tc.input
			if tc.wrongInput != "" {
				testPassword = tc.wrongInput
			}

			if password.Compare(testPassword) != tc.isMatch {
				t.Errorf("[case: %s] password does not match the hash", tc.desc)
				return
			}
		})
	}
}
