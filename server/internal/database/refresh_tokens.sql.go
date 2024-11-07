// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: refresh_tokens.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createRefreshToken = `-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, user_id, created_at, updated_at, expires_at)
VALUES (
    $1,
    $2,
    NOW(),
    NOW(),
    $3
)
RETURNING token, user_id, created_at, updated_at, expires_at, revoked_at
`

type CreateRefreshTokenParams struct {
	Token     string
	UserID    uuid.UUID
	ExpiresAt time.Time
}

func (q *Queries) CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, createRefreshToken, arg.Token, arg.UserID, arg.ExpiresAt)
	var i RefreshToken
	err := row.Scan(
		&i.Token,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
		&i.RevokedAt,
	)
	return i, err
}

const extendRefreshToken = `-- name: ExtendRefreshToken :one
UPDATE refresh_tokens
SET expires_at=$1, updated_at=NOW()
WHERE token=$2
AND revoked_at IS NULL
AND expires_at > NOW()
RETURNING token, user_id, created_at, updated_at, expires_at, revoked_at
`

type ExtendRefreshTokenParams struct {
	ExpiresAt time.Time
	Token     string
}

func (q *Queries) ExtendRefreshToken(ctx context.Context, arg ExtendRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, extendRefreshToken, arg.ExpiresAt, arg.Token)
	var i RefreshToken
	err := row.Scan(
		&i.Token,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
		&i.RevokedAt,
	)
	return i, err
}

const getUserFromRefreshToken = `-- name: GetUserFromRefreshToken :one
SELECT users.id, users.email, users.hashed_password, users.created_at, users.updated_at, users.last_seen_at FROM users
JOIN refresh_tokens ON users.id=refresh_tokens.user_id
WHERE refresh_tokens.token=$1
AND revoked_at IS NULL
AND expires_at > NOW()
`

func (q *Queries) GetUserFromRefreshToken(ctx context.Context, token string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserFromRefreshToken, token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.HashedPassword,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.LastSeenAt,
	)
	return i, err
}

const revokeRefreshToken = `-- name: RevokeRefreshToken :one
UPDATE refresh_tokens
SET revoked_at=NOW(), updated_at=NOW()
WHERE token=$1
AND revoked_at IS NULL
AND expires_at > NOW()
RETURNING token, user_id, created_at, updated_at, expires_at, revoked_at
`

func (q *Queries) RevokeRefreshToken(ctx context.Context, token string) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, revokeRefreshToken, token)
	var i RefreshToken
	err := row.Scan(
		&i.Token,
		&i.UserID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
		&i.RevokedAt,
	)
	return i, err
}
