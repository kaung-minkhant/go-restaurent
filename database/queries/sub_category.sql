-- name: GetAllSubCategory :many
SELECT * FROM menu_sub_categories;

-- name: GetSubCategoryById :one
SELECT * FROM menu_sub_categories
WHERE id = $1;

-- name: CreateSubCategory :one
INSERT INTO menu_sub_categories (id, name, img_url, category, created_at, updated_at, deleted_at)
VALUES ($1, $2, $3, $4, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL)
RETURNING *;

-- name: DeleteSubCategory :one
UPDATE menu_sub_categories
SET deleted_at = NOW() AT TIME ZONE 'utc'
WHERE id = $1
RETURNING *;
