-- name: InsertUser :exec
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
RETURNING *;

-- name: SelectUserByID :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: SelectUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: SelectAllUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
    email = COALESCE(sqlc.narg(email), email),
    password = COALESCE(sqlc.narg(password), password),
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    phone_number = COALESCE(sqlc.narg(phone_number), phone_number),
    role = COALESCE(sqlc.narg(role), role),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;