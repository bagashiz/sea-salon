-- name: InsertUser :one
INSERT INTO users (
    email,
    password,
    full_name,
    phone,
    role
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1
LIMIT 1;

-- name: GetAllUsers :many
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
    phone = COALESCE(sqlc.narg(phone), phone),
    role = COALESCE(sqlc.narg(role), role),
    updated_at = now()
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;