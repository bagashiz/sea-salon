// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: accounts.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const deleteAccount = `-- name: DeleteAccount :one
DELETE FROM accounts
WHERE id = $1
RETURNING id
`

func (q *Queries) DeleteAccount(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, deleteAccount, id)
	err := row.Scan(&id)
	return id, err
}

const insertAccount = `-- name: InsertAccount :exec
INSERT INTO accounts (
    id,
    email,
    password,
    full_name,
    phone_number,
    role,
    created_at,
    updated_at
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
`

type InsertAccountParams struct {
	ID          uuid.UUID
	Email       string
	Password    string
	FullName    string
	PhoneNumber string
	Role        AccountRole
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func (q *Queries) InsertAccount(ctx context.Context, arg InsertAccountParams) error {
	_, err := q.db.Exec(ctx, insertAccount,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.FullName,
		arg.PhoneNumber,
		arg.Role,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	return err
}

const selectAccountByEmail = `-- name: SelectAccountByEmail :one
SELECT id, role, email, password, full_name, phone_number, created_at, updated_at FROM accounts
WHERE email = $1
LIMIT 1
`

func (q *Queries) SelectAccountByEmail(ctx context.Context, email string) (Account, error) {
	row := q.db.QueryRow(ctx, selectAccountByEmail, email)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectAccountByID = `-- name: SelectAccountByID :one
SELECT id, role, email, password, full_name, phone_number, created_at, updated_at FROM accounts
WHERE id = $1
LIMIT 1
`

func (q *Queries) SelectAccountByID(ctx context.Context, id uuid.UUID) (Account, error) {
	row := q.db.QueryRow(ctx, selectAccountByID, id)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectAllAccounts = `-- name: SelectAllAccounts :many
SELECT id, role, email, password, full_name, phone_number, created_at, updated_at FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2
`

type SelectAllAccountsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) SelectAllAccounts(ctx context.Context, arg SelectAllAccountsParams) ([]Account, error) {
	rows, err := q.db.Query(ctx, selectAllAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Role,
			&i.Email,
			&i.Password,
			&i.FullName,
			&i.PhoneNumber,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET
    email = COALESCE($2, email),
    password = COALESCE($3, password),
    full_name = COALESCE($4, full_name),
    phone_number = COALESCE($5, phone_number),
    role = COALESCE($6, role),
    updated_at = now()
WHERE id = $1
RETURNING id, role, email, password, full_name, phone_number, created_at, updated_at
`

type UpdateAccountParams struct {
	ID          uuid.UUID
	Email       pgtype.Text
	Password    pgtype.Text
	FullName    pgtype.Text
	PhoneNumber pgtype.Text
	Role        NullAccountRole
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRow(ctx, updateAccount,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.FullName,
		arg.PhoneNumber,
		arg.Role,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Role,
		&i.Email,
		&i.Password,
		&i.FullName,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}