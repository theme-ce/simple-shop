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

-- name: UpdateProduct :one
UPDATE products
SET
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description),
    price = COALESCE(sqlc.narg(price), price),
    stock_quantity = COALESCE(sqlc.narg(stock_quantity), stock_quantity)
WHERE
    id = sqlc.arg(id)
RETURNING *;