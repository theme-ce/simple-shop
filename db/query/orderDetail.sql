-- name: CreateOrderDetail :one
INSERT INTO "orderDetails" (
    product_id,
    username,
    quantity_ordered,
    price_at_time_of_order
) 
VALUES (
    $1, $2, $3, $4
)
RETURNING *;