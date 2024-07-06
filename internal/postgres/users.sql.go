// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO users (
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
RETURNING id, role, email, password, full_name, phone_number, created_at, updated_at
`

type InsertUserParams struct {
	ID          uuid.UUID
	Email       string
	Password    string
	FullName    string
	PhoneNumber string
	Role        UserRole
	CreatedAt   pgtype.Timestamp
	UpdatedAt   pgtype.Timestamp
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.Exec(ctx, insertUser,
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

const selectAllUsers = `-- name: SelectAllUsers :many
SELECT id, role, email, password, full_name, phone_number, created_at, updated_at FROM users
ORDER BY id
LIMIT $1
OFFSET $2
`

type SelectAllUsersParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) SelectAllUsers(ctx context.Context, arg SelectAllUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, selectAllUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
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

const selectUserByEmail = `-- name: SelectUserByEmail :one
SELECT id, role, email, password, full_name, phone_number, created_at, updated_at FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) SelectUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, selectUserByEmail, email)
	var i User
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

const selectUserByID = `-- name: SelectUserByID :one
SELECT id, role, email, password, full_name, phone_number, created_at, updated_at FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) SelectUserByID(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, selectUserByID, id)
	var i User
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
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

type UpdateUserParams struct {
	ID          uuid.UUID
	Email       pgtype.Text
	Password    pgtype.Text
	FullName    pgtype.Text
	PhoneNumber pgtype.Text
	Role        NullUserRole
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.FullName,
		arg.PhoneNumber,
		arg.Role,
	)
	var i User
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
