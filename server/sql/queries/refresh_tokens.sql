-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, user_id, created_at, updated_at, expires_at)
VALUES (
    $1,
    $2,
    NOW(),
    NOW(),
    $3
)
RETURNING *;

-- name: GetUserFromRefreshToken :one
SELECT users.* FROM users
JOIN refresh_tokens ON users.id=refresh_tokens.user_id
WHERE refresh_tokens.token=$1
AND revoked_at IS NULL
AND expires_at > NOW();

-- name: ExtendRefreshToken :one
UPDATE refresh_tokens
SET expires_at=$1, updated_at=NOW()
WHERE token=$2
RETURNING *;

-- name: RevokeRefreshToken :one
UPDATE refresh_tokens
SET revoked_at=NOW(), updated_at=NOW()
WHERE token=$1
RETURNING *;