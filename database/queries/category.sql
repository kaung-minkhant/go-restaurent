-- name: GetAllCategory :many
SELECT * FROM menu_categories;

-- name: GetCategoryById :one
SELECT * FROM menu_categories
WHERE id = $1;

-- name: CreateCategory :one
INSERT INTO menu_categories (id, name, img_url, created_at, updated_at, deleted_at)
VALUES ($1, $2, $3, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING *;

-- name: DeleteCategory :one
UPDATE menu_categories
SET deleted_at = NOW()
WHERE id = $1
RETURNING *;
