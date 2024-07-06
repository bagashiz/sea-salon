package repository_test

import (
	"context"
	"testing"

	"github.com/bagashiz/sea-salon/internal/app/user"
	"github.com/bagashiz/sea-salon/internal/postgres/repository"
	"github.com/bagashiz/sea-salon/internal/testutil"
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

	err := repo.CreateUser(ctx, u)
	if err != nil {
		t.Errorf("[CreateUser] error: %q", err)
		return
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

	got, err := repo.GetUserByID(ctx, want.ID.String())
	if err != nil {
		t.Errorf("[GetUserByID] error: %q", err)
		return
	}

	if diff := testutil.Diff(got, want); diff != "" {
		t.Error(testutil.Callers(), diff)
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

	got, err := repo.GetUserByEmail(ctx, want.Email)
	if err != nil {
		t.Errorf("[GetUserByEmail] error: %q", err)
		return
	}

	if diff := testutil.Diff(got, want); diff != "" {
		t.Error(testutil.Callers(), diff)
	}

}

func TestListUsers(t *testing.T) {
	ctx := context.Background()
	repo := repository.NewUserRepository(testDB)

	limit := 5
	offset := 0

	for range limit {
		u := createUser()
		err := repo.CreateUser(ctx, u)
		if err != nil {
			t.Errorf("[CreateUser] error: %q", err)
			return
		}
	}

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

	want := createUser()
	err := repo.CreateUser(ctx, want)
	if err != nil {
		t.Errorf("[CreateUser] error: %q", err)
		return
	}

	want.FullName = "Changed Name"
	want.PhoneNumber = "1234567890"

	err = repo.UpdateUser(ctx, want)
	if err != nil {
		t.Errorf("[UpdateUser] error: %q", err)
		return
	}

	got, err := repo.GetUserByID(ctx, want.ID.String())
	if err != nil {
		t.Errorf("[GetUserByID] error: %q", err)
		return
	}

	if diff := testutil.Diff(got, want); diff != "" {
		t.Error(testutil.Callers(), diff)
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

	err = repo.DeleteUser(ctx, want.ID.String())
	if err != nil {
		t.Errorf("[DeleteUser] error: %q", err)
		return
	}

	_, err = repo.GetUserByID(ctx, want.ID.String())
	if err == nil {
		t.Errorf("[GetUserByID] error: %q", err)
		return
	}
}
