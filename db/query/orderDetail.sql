-- name: CreateOrderDetail :one
INSERT INTO "orderDetails" (
    product_id,
    username,
    quantity_ordered,
    price_at_time_of_order,
    order_id
) 
VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;