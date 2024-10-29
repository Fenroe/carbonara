package config

import (
	"time"

	"github.com/google/uuid"
)

type responseUser struct {
	ID         uuid.UUID `json:"id"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	LastSeenAt time.Time `json:"last_seen_at"`
}
