-- +goose Up
CREATE TABLE IF NOT EXISTS role_permissions (
  id SERIAL PRIMARY KEY NOT NULL,
  role VARCHAR(50) NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  permission VARCHAR(50) NOT NULL REFERENCES permissions(permission) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL DEFAULT ( NOW() AT TIME ZONE 'utc' ),
  updated_at TIMESTAMP NOT NULL DEFAULT ( NOW() AT TIME ZONE 'utc' ),
  deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE role_permissions;