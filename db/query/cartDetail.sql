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

-- name: GetCartDetailsByCartID :many
SELECT * FROM "cartDetails"
WHERE cart_id = $1;

-- name: GetCartDetail :one
SELECT * FROM "cartDetails"
WHERE
    cart_id = $1 AND product_id = $2;

-- name: DeleteCartDetail :exec
DELETE FROM "cartDetails"
WHERE
    cart_id = $1 AND product_id = $2;

-- name: ClearCart :exec
DELETE FROM "cartDetails"
WHERE cart_id = $1;