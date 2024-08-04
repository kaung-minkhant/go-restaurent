-- name: CreateRole :one
INSERT INTO roles (id, created_at, updated_at, deleted_at)
VALUES ($1,  NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING *;

-- name: GetAllRoles :many
SELECT * FROM roles;

-- name: GetRoleById :one
SELECT * FROM roles
WHERE id = $1;

-- name: DeleteRole :one
UPDATE roles
SET deleted_at = NOW() AT TIME ZONE 'utc'
WHERE id = $1
RETURNING *;

