-- name: CreatePermission :one
INSERT INTO permissions (id, permission, method, route, created_at, updated_at, deleted_at)
VALUES (DEFAULT, $1, $2, $3, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING *;

-- name: GetPermissionByMethodAndRoute :one
SELECT * FROM permissions
WHERE method = $1
AND route = $2
LIMIT 1;