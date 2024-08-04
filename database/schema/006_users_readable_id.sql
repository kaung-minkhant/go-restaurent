-- +goose Up
ALTER TABLE users 
ADD COLUMN IF NOT EXISTS employee_id VARCHAR(50) UNIQUE NOT NULL DEFAULT ('');

-- +goose Down
ALTER TABLE users
DROP COLUMN employee_id;
