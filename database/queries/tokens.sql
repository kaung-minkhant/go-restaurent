-- name: CreateTokens :one
INSERT INTO auth.tokens (id, access_token, refresh_token, session, valid)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAuthByTokens :one
SELECT * FROM auth.tokens
WHERE refresh_token = $1 AND access_token = $2;

-- name: InvalidateToken :one
UPDATE auth.tokens
SET valid = false
WHERE refresh_token = $1 AND access_token = $2
RETURNING *;

-- name: InvalidateTokenFamily :exec
DELETE FROM auth.tokens
WHERE session = $1;