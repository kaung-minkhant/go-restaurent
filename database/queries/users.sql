-- name: CreateUser :one
INSERT INTO users (id, name, encrypted_password, role, created_at, updated_at, deleted_at, employee_id)
VALUES ($1, $2, $3, $4, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL, $5)
RETURNING *;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserByEmployeeId :one
SELECT * FROM users
WHERE employee_id = $1;

-- name: DeleteUser :one
UPDATE users
SET deleted_at = NOW() AT TIME ZONE 'utc'
WHERE id = $1
RETURNING *;

