-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    first_name,
    last_name,
    email
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;