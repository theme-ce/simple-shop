-- name: CreateCart :one
INSERT INTO carts (
    username
) VALUES (
    $1
)
RETURNING *;
