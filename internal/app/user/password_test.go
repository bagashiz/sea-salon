package user_test

import (
	"fmt"
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

// TestPassword tests the NewPassword and Compare methods of the Password type.
func TestPassword(t *testing.T) {
	testCases := []struct {
		err       error
		desc      string
		input     string
		wrongHash user.Password
		isMatch   bool
	}{
		{desc: "valid_password_0", input: "test1234", wrongHash: "", isMatch: true, err: nil},
		{desc: "valid_password_1", input: "#tes+t1@-4e2=34$", wrongHash: "", isMatch: true, err: nil},
		{desc: "valid_password_2", input: "$2a$10$4tS9MQtS6l/9PWY.MiR8O.3.yFKHvaB34kpQVGgSVnGrla6ztOaam", wrongHash: "", isMatch: true, err: nil},
		{desc: "empty_password", input: "", wrongHash: "", isMatch: true, err: user.ErrPasswordEmpty},
		{desc: "password_too_short", input: "test", wrongHash: "", isMatch: true, err: user.ErrPasswordTooShort},
		{desc: "password_too_long", input: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", wrongHash: "", isMatch: true, err: user.ErrPasswordTooLong},
		{desc: "invalid_password_0", input: "test\t1234", wrongHash: "", isMatch: true, err: user.ErrPasswordInvalid},
		{desc: "invalid_password_1", input: " test1234 ", wrongHash: "", isMatch: true, err: user.ErrPasswordInvalid},
		{desc: "different_password", input: "test1234", wrongHash: "$2a$10$4tS9MQtS6l/9PWY.MiR8O.3.yFKHvaB34kpQVGgSVnGrla6ztOaam", isMatch: false, err: nil},
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

			hashedPassword, err := password.Hash()
			if err != nil {
				t.Errorf("[case: %s] failed to hash the password", tc.desc)
				return
			}

			if tc.wrongHash != "" {
				hashedPassword = tc.wrongHash
			}

			fmt.Println(hashedPassword)
			fmt.Println(tc.wrongHash)

			if password.Compare(hashedPassword.String()) != tc.isMatch {
				t.Errorf("[case: %s] password does not match the hash", tc.desc)
				return
			}
		})
	}
}
