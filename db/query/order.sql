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

-- name: UpdateOrderStatus :one
UPDATE orders
SET
    status = $1
WHERE
    id = $2
RETURNING *;