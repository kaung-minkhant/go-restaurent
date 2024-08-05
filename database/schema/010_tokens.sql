-- +goose Up
CREATE TABLE IF NOT EXISTS auth.tokens (
  id UUID PRIMARY KEY NOT NULL UNIQUE,
  refresh_token TEXT NOT NULL UNIQUE,
  session UUID NOT NULL,
  access_token TEXT NOT NULL UNIQUE,
  valid BOOLEAN NOT NULL DEFAULT FALSE
);

-- +goose Down
DROP TABLE auth.tokens;