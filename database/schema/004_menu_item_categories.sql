-- +goose Up
ALTER TABLE menu_items
ADD COLUMN category UUID REFERENCES menu_categories(id) ON DELETE SET NULL,
ADD COLUMN sub_category UUID REFERENCES menu_sub_categories(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE menu_items
DROP COLUMN category,
DROP COLUMN sub_category;