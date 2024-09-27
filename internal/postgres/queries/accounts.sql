-- name: InsertAccount :exec
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
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: SelectAccountByID :one
SELECT * FROM accounts
WHERE id = $1
LIMIT 1;

-- name: SelectAccountByEmail :one
SELECT * FROM accounts
WHERE email = $1
LIMIT 1;

-- name: SelectAllAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET
    email = COALESCE(sqlc.narg(email), email),
    password = COALESCE(sqlc.narg(password), password),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    phone_number = COALESCE(sqlc.narg(phone_number), phone_number),
    role = COALESCE(sqlc.narg(role), role),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteAccount :one
DELETE FROM accounts
WHERE id = $1
RETURNING id;
