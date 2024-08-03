-- +goose Up
CREATE TABLE IF NOT EXISTS menu_sub_categories (
  id UUID NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  img_url TEXT NOT NULL,
  category UUID NOT NULL REFERENCES menu_categories(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE menu_sub_categories;