// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type RevokedToken struct {
	ID        uuid.UUID
	Token     string
	RevokedAt sql.NullTime
}

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}