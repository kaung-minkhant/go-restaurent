-- name: CreateRolePermission :one
INSERT INTO role_permissions (id, role, permission, created_at, updated_at, deleted_at)
VALUES (DEFAULT, $1, $2, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING *;

-- name: GetRoleByPermissionName :one
SELECT * FROM role_permissions
WHERE permission = $1
LIMIT 1;