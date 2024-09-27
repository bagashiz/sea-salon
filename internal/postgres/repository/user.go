package repository

import (
	"context"

	"github.com/bagashiz/sea-salon/internal/app/user"
	"github.com/bagashiz/sea-salon/internal/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

// AddAccount inserts user account data to postgres database.
func (r *PostgresRepository) AddAccount(ctx context.Context, account *user.Account) error {
	arg := postgres.InsertAccountParams{
		ID:          account.ID,
		Email:       account.Email,
		Password:    account.Password,
		FullName:    account.FullName,
		PhoneNumber: account.PhoneNumber,
		Role:        postgres.AccountRole(account.Role),
		CreatedAt:   pgtype.Timestamptz{Time: account.CreatedAt, Valid: !account.CreatedAt.IsZero()},
		UpdatedAt:   pgtype.Timestamptz{Time: account.UpdatedAt, Valid: !account.UpdatedAt.IsZero()},
	}

	if err := r.db.InsertAccount(ctx, arg); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgerrcode.IsDataException(pgErr.Code) {
				return user.ErrAccountInvalid
			}
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return user.ErrAccountExists
			}
		}
		return err
	}

	return nil
}

// GetAccountByID retrieves user account data from postgres database by ID.
func (r *PostgresRepository) GetAccountByID(ctx context.Context, id uuid.UUID) (*user.Account, error) {
	result, err := r.db.SelectAccountByID(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user.ErrAccountNotFound
		}
		return nil, err
	}

	account := result.ToDomain()

	return account, nil
}

// GetAccountByEmail retrieves user account data from postgres database by email.
func (r *PostgresRepository) GetAccountByEmail(ctx context.Context, email string) (*user.Account, error) {
	result, err := r.db.SelectAccountByEmail(ctx, email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, user.ErrAccountNotFound
		}
		return nil, err
	}

	account := result.ToDomain()

	return account, nil
}

// ListAccounts retrieves a list of user accounts from postgres database.
func (r *PostgresRepository) ListAccounts(ctx context.Context, limit, offset int) ([]*user.Account, error) {
	arg := postgres.SelectAllAccountsParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	}

	result, err := r.db.SelectAllAccounts(ctx, arg)
	if err != nil {
		return nil, err
	}

	var accounts []*user.Account

	for _, account := range result {
		accounts = append(accounts, account.ToDomain())
	}

	return accounts, nil
}

// UpdateAccount updates user account data in postgres database.
func (r *PostgresRepository) UpdateAccount(ctx context.Context, account *user.Account) error {
	arg := postgres.UpdateAccountParams{
		ID:          account.ID,
		Email:       pgtype.Text{String: account.Email, Valid: account.Email != ""},
		Password:    pgtype.Text{String: account.Password, Valid: account.Password != ""},
		FullName:    pgtype.Text{String: account.FullName, Valid: account.FullName != ""},
		PhoneNumber: pgtype.Text{String: account.PhoneNumber, Valid: account.PhoneNumber != ""},
		Role:        postgres.NullAccountRole{AccountRole: postgres.AccountRole(account.Role), Valid: account.Role != ""},
	}

	result, err := r.db.UpdateAccount(ctx, arg)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			if pgerrcode.IsDataException(pgErr.Code) {
				return user.ErrAccountInvalid
			}
			if pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
				return user.ErrAccountExists
			}
		}
		if err == pgx.ErrNoRows {
			return user.ErrAccountNotFound
		}
		return err
	}

	account = result.ToDomain()
	_ = account // avoid linter warning

	return nil
}

// DeleteAccount removes user data from postgres database.
func (r *PostgresRepository) DeleteAccount(ctx context.Context, id uuid.UUID) error {
	_, err := r.db.DeleteAccount(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return user.ErrAccountNotFound
		}
		return err
	}

	return nil
}
