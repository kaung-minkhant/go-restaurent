
-- +goose Up
CREATE TABLE IF NOT EXISTS menu_categories (
  id UUID NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  img_url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE menu_categories;