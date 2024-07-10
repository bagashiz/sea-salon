package repository_test

import (
	"context"
	"sync"
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
	"github.com/bagashiz/sea-salon/internal/postgres/repository"
	"github.com/bagashiz/sea-salon/internal/testutil"
	"github.com/google/uuid"
)

// createUser creates a new user with random data.
func createUser() *user.User {
	fullName := testutil.RandomFullName()
	phoneNumber := testutil.RandomPhoneNumber()
	role := "customer"
	email := testutil.RandomEmail()
	password, err := user.NewPassword(testutil.RandomAlphaNumeric(8))
	if err != nil {
		return nil
	}
	hashedPassword, err := password.Hash()
	if err != nil {
		return nil
	}

	u := &user.User{
		Email:       email,
		Password:    hashedPassword,
		FullName:    fullName,
		PhoneNumber: phoneNumber,
		Role:        role,
	}
	u.Create()

	return u
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewUserRepository(testDB)
	u := createUser()

	testCases := []struct {
		desc string
		user *user.User
		err  error
	}{
		{desc: "valid user", user: u, err: nil},
		{desc: "invalid user", user: &user.User{}, err: user.ErrUserInvalid},
		{desc: "invalid role", user: &user.User{Role: "invalidrole"}, err: user.ErrUserInvalid},
		{desc: "duplicate user", user: u, err: user.ErrUserExists},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := repo.CreateUser(ctx, tc.user)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewUserRepository(testDB)

	want := createUser()
	err := repo.CreateUser(ctx, want)
	if err != nil {
		t.Errorf("[CreateUser] error: %q", err)
		return
	}

	testCases := []struct {
		desc string
		id   string
		err  error
	}{
		{desc: "existing user ID", id: want.ID.String(), err: nil},
		{desc: "invalid user ID", id: "invaliduuid", err: user.ErrIDInvalid},
		{desc: "non-existing user ID", id: uuid.NewString(), err: user.ErrUserNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := repo.GetUserByID(ctx, tc.id)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}

			if err == nil {
				if diff := testutil.Diff(got, want); diff != "" {
					t.Errorf("[case: %s] %s %s", tc.desc, testutil.Callers(), diff)
				}
			}
		})
	}
}

func TestGetUserByEmail(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewUserRepository(testDB)

	want := createUser()
	err := repo.CreateUser(ctx, want)
	if err != nil {
		t.Errorf("[CreateUser] error: %q", err)
		return
	}

	testCases := []struct {
		desc  string
		email string
		err   error
	}{
		{desc: "existing email", email: want.Email, err: nil},
		{desc: "non-existing email", email: "notexists@email.com", err: user.ErrUserNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := repo.GetUserByEmail(ctx, tc.email)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}

			if err == nil {
				if diff := testutil.Diff(got, want); diff != "" {
					t.Errorf("[case: %s] %s %s", tc.desc, testutil.Callers(), diff)
				}
			}
		})
	}
}

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewUserRepository(testDB)

	if _, err := testDB.Exec(ctx, "DELETE FROM users"); err != nil {
		t.Errorf("[ListUsers] error: %q", err)
		return
	}

	limit := 5
	offset := 0

	var wg sync.WaitGroup

	for range limit {
		wg.Add(1)
		go func() {
			defer wg.Done()
			u := createUser()
			err := repo.CreateUser(ctx, u)
			if err != nil {
				t.Errorf("[CreateUser] error: %q", err)
				return
			}
		}()
	}

	wg.Wait()

	got, err := repo.ListUsers(ctx, limit, offset)
	if err != nil {
		t.Errorf("[ListUsers] error: %q", err)
		return
	}

	if len(got) != limit {
		t.Errorf("expected %d users, got %d", limit, len(got))
		return
	}
}

func TestUpdateUser(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewUserRepository(testDB)

	dummy1 := createUser()
	err := repo.CreateUser(ctx, dummy1)
	if err != nil {
		t.Errorf("[CreateUser] error: %q", err)
		return
	}

	dummy2 := createUser()
	err = repo.CreateUser(ctx, dummy2)
	if err != nil {
		t.Errorf("[CreateUser] error: %q", err)
		return
	}

	testCases := []struct {
		desc string
		user *user.User
		err  error
	}{
		{desc: "valid user", user: dummy1, err: nil},
		{desc: "update email", user: &user.User{ID: dummy1.ID, Email: testutil.RandomEmail()}, err: nil},
		{desc: "update full name", user: &user.User{ID: dummy1.ID, FullName: testutil.RandomFullName()}, err: nil},
		{desc: "update phone number", user: &user.User{ID: dummy1.ID, PhoneNumber: testutil.RandomPhoneNumber()}, err: nil},
		{desc: "update role", user: &user.User{ID: dummy1.ID, Role: "admin"}, err: nil},
		{desc: "invalid user", user: &user.User{}, err: user.ErrUserNotFound},
		{desc: "non-existing user", user: &user.User{ID: uuid.New()}, err: user.ErrUserNotFound},
		{desc: "duplicate email", user: &user.User{ID: dummy1.ID, Email: dummy2.Email}, err: user.ErrUserExists},
		{desc: "invalid role", user: &user.User{ID: dummy1.ID, Role: "invalidrole"}, err: user.ErrUserInvalid},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := repo.UpdateUser(ctx, tc.user)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewUserRepository(testDB)

	want := createUser()
	err := repo.CreateUser(ctx, want)
	if err != nil {
		t.Errorf("[CreateUser] error: %q", err)
		return
	}

	testCases := []struct {
		desc string
		id   string
		err  error
	}{
		{desc: "existing user ID", id: want.ID.String(), err: nil},
		{desc: "invalid user ID", id: "invaliduuid", err: user.ErrIDInvalid},
		{desc: "non-existing user ID", id: uuid.NewString(), err: user.ErrUserNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := repo.DeleteUser(ctx, tc.id)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}
		})
	}
}
