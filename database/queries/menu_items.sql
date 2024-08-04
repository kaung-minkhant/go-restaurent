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
UPDATE menu_items
SET deleted_at = NOW() AT TIME ZONE 'utc'
WHERE id = $1
RETURNING *;


-- name: HardDeleteMenuItem :one
DELETE FROM menu_items
WHERE id = $1
RETURNING *;
