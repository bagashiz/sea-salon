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

// TestAddAccount tests the AddAccount method.
func TestAddAccount(t *testing.T) {
	ctx := context.Background()
	repo := repository.New(testDB)
	account, err := createAccount()
	if err != nil {
		t.Errorf("[createAccount] error: %q", err)
		return
	}

	testCases := []struct {
		err     error
		account *user.Account
		desc    string
	}{
		{desc: "valid_account", account: account, err: nil},
		{desc: "invalid_account", account: &user.Account{}, err: user.ErrAccountInvalid},
		{desc: "invalid_role", account: &user.Account{Role: "invalidrole"}, err: user.ErrAccountInvalid},
		{desc: "duplicate_account", account: account, err: user.ErrAccountExists},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := repo.AddAccount(ctx, tc.account)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}
		})
	}
}

// TestGetAccountByID tests the GetAccountByID method.
func TestGetAccountByID(t *testing.T) {
	ctx := context.Background()
	repo := repository.New(testDB)

	want, err := addDummyAccount(ctx, repo)
	if err != nil {
		t.Errorf("[addDummyAccount] error: %q", err)
	}

	testCases := []struct {
		err  error
		desc string
		id   uuid.UUID
	}{
		{desc: "existing_account_ID", id: want.ID, err: nil},
		{desc: "non-existing_account_ID", id: uuid.New(), err: user.ErrAccountNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := repo.GetAccountByID(ctx, tc.id)
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

// TestGetAccountByEmail tests the GetAccountByEmail method.
func TestGetAccountByEmail(t *testing.T) {
	ctx := context.Background()
	repo := repository.New(testDB)

	want, err := addDummyAccount(ctx, repo)
	if err != nil {
		t.Errorf("[addDummyAccount] error: %q", err)
	}

	testCases := []struct {
		err   error
		desc  string
		email string
	}{
		{desc: "existing_email", email: want.Email, err: nil},
		{desc: "non-existing_email", email: "notexists@email.com", err: user.ErrAccountNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got, err := repo.GetAccountByEmail(ctx, tc.email)
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

// TestListAccounts tests the ListAccounts method.
func TestListAccounts(t *testing.T) {
	ctx := context.Background()
	repo := repository.New(testDB)

	if _, err := testDB.Exec(ctx, "DELETE FROM accounts"); err != nil {
		t.Errorf("[ListAccounts] error: %q", err)
		return
	}

	limit := 5
	offset := 0

	var wg sync.WaitGroup

	for range limit {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := addDummyAccount(ctx, repo)
			if err != nil {
				t.Errorf("[addDummyAccount] error: %q", err)
			}
		}()
	}

	wg.Wait()

	got, err := repo.ListAccounts(ctx, limit, offset)
	if err != nil {
		t.Errorf("[ListAccounts] error: %q", err)
		return
	}

	if len(got) != limit {
		t.Errorf("expected %d accounts, got %d", limit, len(got))
		return
	}
}

// TestUpdateAccount tests the UpdateAccount method.
func TestUpdateAccount(t *testing.T) {
	ctx := context.Background()
	repo := repository.New(testDB)

	dummy1, err := addDummyAccount(ctx, repo)
	if err != nil {
		t.Errorf("[addDummyAccount] error: %q", err)
	}

	dummy2, err := addDummyAccount(ctx, repo)
	if err != nil {
		t.Errorf("[addDummyAccount] error: %q", err)
	}

	testCases := []struct {
		err     error
		account *user.Account
		desc    string
	}{
		{desc: "valid_account", account: dummy1, err: nil},
		{desc: "update_email", account: &user.Account{ID: dummy1.ID, Email: testutil.RandomEmail()}, err: nil},
		{desc: "update_full_name", account: &user.Account{ID: dummy1.ID, FullName: testutil.RandomFullName()}, err: nil},
		{desc: "update_phone_number", account: &user.Account{ID: dummy1.ID, PhoneNumber: testutil.RandomPhoneNumber()}, err: nil},
		{desc: "update_role", account: &user.Account{ID: dummy1.ID, Role: "admin"}, err: nil},
		{desc: "invalid_account", account: &user.Account{}, err: user.ErrAccountNotFound},
		{desc: "non-existing_account", account: &user.Account{ID: uuid.New()}, err: user.ErrAccountNotFound},
		{desc: "duplicate_email", account: &user.Account{ID: dummy1.ID, Email: dummy2.Email}, err: user.ErrAccountExists},
		{desc: "invalid_role", account: &user.Account{ID: dummy1.ID, Role: "invalidrole"}, err: user.ErrAccountInvalid},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := repo.UpdateAccount(ctx, tc.account)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}
		})
	}
}

// TestDeleteAccount tests the DeleteAccount method.
func TestDeleteAccount(t *testing.T) {
	ctx := context.Background()
	repo := repository.New(testDB)

	want, err := addDummyAccount(ctx, repo)
	if err != nil {
		t.Errorf("[addDummyAccount] error: %q", err)
	}

	testCases := []struct {
		err  error
		desc string
		id   uuid.UUID
	}{
		{desc: "existing_account_ID", id: want.ID, err: nil},
		{desc: "non-existing_account_ID", id: uuid.New(), err: user.ErrAccountNotFound},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := repo.DeleteAccount(ctx, tc.id)
			if err != tc.err {
				t.Errorf("[case: %s] want %q, got %q", tc.desc, tc.err, err)
				return
			}
		})
	}
}

// createAccount creates a new user account struct with random values.
func createAccount() (*user.Account, error) {
	fullName, err := user.NewFullName(testutil.RandomFullName())
	if err != nil {
		return nil, err
	}

	phoneNumber, err := user.NewPhoneNumber(testutil.RandomPhoneNumber())
	if err != nil {
		return nil, err
	}

	role, err := user.NewAccountRole("Customer")
	if err != nil {
		return nil, err
	}

	email, err := user.NewEmail(testutil.RandomEmail())
	if err != nil {
		return nil, err
	}

	password, err := user.NewPassword(testutil.RandomAlphaNumeric(8))
	if err != nil {
		return nil, err
	}

	hashedPassword, err := password.Hash()
	if err != nil {
		return nil, err
	}

	account, err := user.NewAccount(fullName, phoneNumber, email, hashedPassword, role)
	if err != nil {
		return nil, err
	}

	return account, nil
}

// addDummyAccount inserts a dummy user account data to the database.
func addDummyAccount(ctx context.Context, repo *repository.PostgresRepository) (*user.Account, error) {
	want, err := createAccount()
	if err != nil {
		return nil, err
	}
	err = repo.AddAccount(ctx, want)
	if err != nil {
		return nil, err
	}
	return want, nil
}
