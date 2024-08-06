-- name: CreateTokens :one
INSERT INTO auth.tokens (id, access_token, refresh_token, session, valid)
VALUES ($1, $2, $3, $4, $5) RETURNING *;

-- name: GetAuthByTokens :one
SELECT * FROM auth.tokens
WHERE refresh_token = $1 AND access_token = $2;

-- name: InvalidateToken :one
UPDATE auth.tokens
SET valid = false
WHERE refresh_token = $1 AND access_token = $2
RETURNING *;

-- TODO: Change invalidation to setting all to invalid, then when user try to refresh the token, if his current one is invalidated, then someone has used old tokens to request new token
-- name: InvalidateTokenFamily :exec
UPDATE auth.tokens
SET valid = FALSE
WHERE session IN (
  SELECT session FROM auth.tokens token
  WHERE token.access_token = $1 AND token.refresh_token = $2
) AND valid = TRUE;


-- name: LogoutTokens :exec
DELETE FROM auth.tokens
WHERE session IN 
( 
  SELECT session FROM auth.tokens tokens
  WHERE tokens.access_token = $1 AND tokens.refresh_token = $2
);