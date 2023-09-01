-- name: CreateCartDetail :one
INSERT INTO "cartDetails" (
  cart_id,
  product_id,
  quantity_added
) VALUES (
    $1, $2, $3
)
RETURNING *;