-- name: CreateClip :one
INSERT INTO clips (id, content, user_id, created_at, updated_at)
VALUES (
    GEN_RANDOM_UUID(),
    $1,
    $2,
    NOW(),
    NOW()
)
RETURNING *;

-- name: GetMostRecentClip :one
SELECT * FROM clips
WHERE id IN (
    SELECT * FROM clips c
    WHERE c.user_id=$1
    ORDER BY created_at DESC
    LIMIT 1
);

-- name: GetClipsByUser :many
SELECT * FROM clips
WHERE user_id=$1
ORDER BY created_at DESC;

-- name: DeleteOldestClip :exec
DELETE FROM clips WHERE id IN (
    SELECT * FROM clips c
    WHERE c.user_id=$1
    ORDER BY created_at ASC
    LIMIT 1
);