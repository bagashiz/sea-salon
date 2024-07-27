package user_test

import (
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

func TestNewRole(t *testing.T) {
	testCases := []struct {
		err   error
		desc  string
		input string
		want  user.UserRole
	}{
		{desc: "admin", input: "admin", want: user.Admin, err: nil},
		{desc: "customer", input: "customer", want: user.Customer, err: nil},
		{desc: "invalid", input: "invalid", want: "", err: user.ErrRoleInvalid},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := user.NewRole(tc.input)
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
