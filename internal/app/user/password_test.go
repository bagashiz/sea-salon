package user_test

import (
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

// TestNewPassword tests the NewPassword function.
func TestNewPassword(t *testing.T) {
	testCases := []struct {
		desc  string
		input string
		want  user.Password
		err   error
	}{
		{desc: "valid password", input: "test1234", want: user.Password("test1234"), err: nil},
		{desc: "valid password", input: "#tes+t1@-4e2=34$", want: user.Password("#tes+t1@-4e2=34$"), err: nil},
		{desc: "valid password", input: "$2a$10$4tS9MQtS6l/9PWY.MiR8O.3.yFKHvaB34kpQVGgSVnGrla6ztOaam", want: user.Password("$2a$10$4tS9MQtS6l/9PWY.MiR8O.3.yFKHvaB34kpQVGgSVnGrla6ztOaam"), err: nil},
		{desc: "empty password", input: "", want: "", err: user.ErrPasswordEmpty},
		{desc: "password too short", input: "test", want: "", err: user.ErrPasswordTooShort},
		{desc: "password too long", input: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", want: "", err: user.ErrPasswordTooLong},
		{desc: "invalid password", input: "test\t1234", want: "", err: user.ErrPasswordInvalid},
		{desc: "invalid password", input: " test1234 ", want: "", err: user.ErrPasswordInvalid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := user.NewPassword(tc.input)
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

// TestHashAndComparePassword tests the Hash and Compare methods of the Password type.
func TestHashAndComparePassword(t *testing.T) {
	testCases := []struct {
		desc      string
		password  user.Password
		wrongHash string
		isMatch   bool
	}{
		{desc: "valid password", password: user.Password("test1234"), wrongHash: "", isMatch: true},
		{desc: "valid password", password: user.Password("#tes+t1@-4e2=34$"), wrongHash: "", isMatch: true},
		{desc: "invalid password", password: user.Password("test1234"), wrongHash: "$2a$10$", isMatch: false},
		{desc: "invalid password", password: user.Password("#tes+t1@-4e2=34$"), wrongHash: "$2a$10$", isMatch: false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			hash, err := tc.password.Hash()
			if err != nil {
				t.Errorf("[case: %s] fail to hash password: %q", tc.desc, err)
				return
			}

			if tc.wrongHash != "" {
				hash = tc.wrongHash
			}

			if tc.password.Compare(hash) != tc.isMatch {
				t.Errorf("[case: %s] password does not match the hash", tc.desc)
				return
			}
		})
	}
}
