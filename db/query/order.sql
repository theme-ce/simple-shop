-- name: CreateOrder :one
INSERT INTO orders (
    username,
    total_price,
    status
) 
VALUES (
    $1, $2, $3
)
RETURNING *;