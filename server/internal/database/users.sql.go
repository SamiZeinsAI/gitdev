// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
        id,
        access_token,
        name,
        username,
        github_id,
        repos,
        email,
        bio,
        avatar_url
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) ON CONFLICT (github_id) DO
UPDATE
SET access_token = $2,
    updated_at = $10
RETURNING id, created_at, updated_at, access_token, name, username, github_id, repos, email, bio, avatar_url
`

type CreateUserParams struct {
	ID          uuid.UUID
	AccessToken string
	Name        string
	Username    string
	GithubID    int32
	Repos       int32
	Email       string
	Bio         string
	AvatarUrl   string
	UpdatedAt   time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.AccessToken,
		arg.Name,
		arg.Username,
		arg.GithubID,
		arg.Repos,
		arg.Email,
		arg.Bio,
		arg.AvatarUrl,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessToken,
		&i.Name,
		&i.Username,
		&i.GithubID,
		&i.Repos,
		&i.Email,
		&i.Bio,
		&i.AvatarUrl,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, created_at, updated_at, access_token, name, username, github_id, repos, email, bio, avatar_url
FROM users
LIMIT 20
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AccessToken,
			&i.Name,
			&i.Username,
			&i.GithubID,
			&i.Repos,
			&i.Email,
			&i.Bio,
			&i.AvatarUrl,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByToken = `-- name: GetUserByToken :one
SELECT id, created_at, updated_at, access_token, name, username, github_id, repos, email, bio, avatar_url
FROM users
WHERE github_id = $1
`

func (q *Queries) GetUserByToken(ctx context.Context, githubID int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByToken, githubID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AccessToken,
		&i.Name,
		&i.Username,
		&i.GithubID,
		&i.Repos,
		&i.Email,
		&i.Bio,
		&i.AvatarUrl,
	)
	return i, err
}
