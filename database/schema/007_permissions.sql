-- +goose Up
CREATE TABLE IF NOT EXISTS permissions (
  id SERIAL PRIMARY KEY,
  permission VARCHAR(50) UNIQUE NOT NULL,
  method VARCHAR(10) NOT NULL,
  route TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT ( NOW() AT TIME ZONE 'utc' ),
  updated_at TIMESTAMP NOT NULL DEFAULT ( NOW() AT TIME ZONE 'utc' ),
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE permissions;