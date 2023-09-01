-- name: CreateCart :one
INSERT INTO carts (
    username
) VALUES (
    $1
)
RETURNING *;

-- name: GetCartByID :one
SELECT * FROM carts
WHERE id = $1;

-- name: GetCartByUsername :one
SELECT * FROM carts
WHERE username = $1;

-- name: DeleteCart :exec
DELETE FROM carts
WHERE username = $1;