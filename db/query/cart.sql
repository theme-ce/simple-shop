-- name: CreateCart :one
INSERT INTO carts (
    username
) VALUES (
    $1
)
RETURNING *;

-- name: GetCart :one
SELECT * FROM carts
WHERE username = $1;

-- name: DeleteCart :exec
DELETE FROM carts
WHERE username = $1;