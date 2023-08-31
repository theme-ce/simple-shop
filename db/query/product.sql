-- name: CreateProduct :one
INSERT INTO products (
    name,
    description,
    price,
    stock_quantity
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;