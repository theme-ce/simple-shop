-- name: CreateCartDetail :one
INSERT INTO "cartDetails" (
    cart_id,
    product_id,
    quantity_added
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateCartDetail :one
UPDATE "cartDetails"
SET
    quantity_added = $1
WHERE
    cart_id = $2 AND product_id = $3
RETURNING *;
