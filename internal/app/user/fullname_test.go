package user_test

import (
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
)

// TestNewFullName tests the NewFullName function.
func TestNewFullName(t *testing.T) {
	testCases := []struct {
		err   error
		desc  string
		input string
		want  user.FullName
	}{
		{desc: "valid_fullname", input: "John Doe", want: user.FullName("John Doe"), err: nil},
		{desc: "contains_whitespace", input: "\tJohn Doe ", want: user.FullName("John Doe"), err: nil},
		{desc: "empty_fullname", input: "", want: user.FullName(""), err: user.ErrFullNameEmpty},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			t.Parallel()
			got, err := user.NewFullName(tc.input)
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
