-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
  id VARCHAR(50) PRIMARY KEY NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
  id UUID PRIMARY KEY NOT NULL,
  name TEXT NOT NULL,
  encrypted_password TEXT NOT NULL,
  role VARCHAR(50) NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  deleted_at TIMESTAMP
);


-- +goose Down
DROP TABLE users;
DROP TABLE roles;