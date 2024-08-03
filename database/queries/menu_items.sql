-- name: GetAllMenuItems :many
SELECT * FROM menu_items;

-- name: GetMenuItemById :one
SELECT * FROM menu_items
WHERE id = $1;

-- name: CreateMenuItem :one
INSERT INTO menu_items (id, name, price, description, ingredients, img_url, created_at, updated_at, deleted_at, category, sub_category)
VALUES ($1, $2, $3, $4, $5, $6, NOW() AT TIME ZONE 'utc', NOW() AT TIME ZONE 'utc', NULL, $7, $8)
RETURNING *;

-- name: DeleteMenuItem :one
DELETE FROM menu_items
WHERE id = $1
RETURNING *;
