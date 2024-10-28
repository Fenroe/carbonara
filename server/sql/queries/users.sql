-- name: CreateUser :one
INSERT INTO users (id, email, hashed_password, created_at, updated_at, last_seen_at)
VALUES (
    GEN_RANDOM_UUID(),
    $1,
    $2,
    NOW(),
    NOW(),
    NOW()
)
-- don't return hashed_password
RETURNING id, email, created_at, updated_at, last_seen_at;