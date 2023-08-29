// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    first_name,
    last_name,
    email
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING username, email, hashed_password, first_name, last_name, address, password_changed_at, created_at, is_admin
`

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.HashedPassword,
		arg.FirstName,
		arg.LastName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Address,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsAdmin,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, email, hashed_password, first_name, last_name, address, password_changed_at, created_at, is_admin FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.HashedPassword,
		&i.FirstName,
		&i.LastName,
		&i.Address,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.IsAdmin,
	)
	return i, err
}
