-- +goose Up
CREATE TABLE IF NOT EXISTS menu_items (
  id UUID NOT NULL PRIMARY KEY,
  name TEXT NOT NULL,
  price REAL NOT NULL,
  description TEXT NOT NULL,
  ingredients TEXT NOT NULL,
  img_url TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE menu_items;