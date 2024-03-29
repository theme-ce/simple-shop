// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: cartDetail.sql

package db

import (
	"context"
)

const clearCart = `-- name: ClearCart :exec
DELETE FROM "cartDetails"
WHERE cart_id = $1
`

func (q *Queries) ClearCart(ctx context.Context, cartID int64) error {
	_, err := q.db.Exec(ctx, clearCart, cartID)
	return err
}

const createCartDetail = `-- name: CreateCartDetail :one
INSERT INTO "cartDetails" (
    cart_id,
    product_id,
    quantity_added
) VALUES (
    $1, $2, $3
)
RETURNING id, cart_id, product_id, quantity_added
`

type CreateCartDetailParams struct {
	CartID        int64 `json:"cart_id"`
	ProductID     int64 `json:"product_id"`
	QuantityAdded int64 `json:"quantity_added"`
}

func (q *Queries) CreateCartDetail(ctx context.Context, arg CreateCartDetailParams) (CartDetail, error) {
	row := q.db.QueryRow(ctx, createCartDetail, arg.CartID, arg.ProductID, arg.QuantityAdded)
	var i CartDetail
	err := row.Scan(
		&i.ID,
		&i.CartID,
		&i.ProductID,
		&i.QuantityAdded,
	)
	return i, err
}

const deleteCartDetail = `-- name: DeleteCartDetail :exec
DELETE FROM "cartDetails"
WHERE
    cart_id = $1 AND product_id = $2
`

type DeleteCartDetailParams struct {
	CartID    int64 `json:"cart_id"`
	ProductID int64 `json:"product_id"`
}

func (q *Queries) DeleteCartDetail(ctx context.Context, arg DeleteCartDetailParams) error {
	_, err := q.db.Exec(ctx, deleteCartDetail, arg.CartID, arg.ProductID)
	return err
}

const getCartDetail = `-- name: GetCartDetail :one
SELECT id, cart_id, product_id, quantity_added FROM "cartDetails"
WHERE
    cart_id = $1 AND product_id = $2
`

type GetCartDetailParams struct {
	CartID    int64 `json:"cart_id"`
	ProductID int64 `json:"product_id"`
}

func (q *Queries) GetCartDetail(ctx context.Context, arg GetCartDetailParams) (CartDetail, error) {
	row := q.db.QueryRow(ctx, getCartDetail, arg.CartID, arg.ProductID)
	var i CartDetail
	err := row.Scan(
		&i.ID,
		&i.CartID,
		&i.ProductID,
		&i.QuantityAdded,
	)
	return i, err
}

const getCartDetailsByCartID = `-- name: GetCartDetailsByCartID :many
SELECT id, cart_id, product_id, quantity_added FROM "cartDetails"
WHERE cart_id = $1
`

func (q *Queries) GetCartDetailsByCartID(ctx context.Context, cartID int64) ([]CartDetail, error) {
	rows, err := q.db.Query(ctx, getCartDetailsByCartID, cartID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CartDetail{}
	for rows.Next() {
		var i CartDetail
		if err := rows.Scan(
			&i.ID,
			&i.CartID,
			&i.ProductID,
			&i.QuantityAdded,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCartDetail = `-- name: UpdateCartDetail :one
UPDATE "cartDetails"
SET
    quantity_added = $1
WHERE
    cart_id = $2 AND product_id = $3
RETURNING id, cart_id, product_id, quantity_added
`

type UpdateCartDetailParams struct {
	QuantityAdded int64 `json:"quantity_added"`
	CartID        int64 `json:"cart_id"`
	ProductID     int64 `json:"product_id"`
}

func (q *Queries) UpdateCartDetail(ctx context.Context, arg UpdateCartDetailParams) (CartDetail, error) {
	row := q.db.QueryRow(ctx, updateCartDetail, arg.QuantityAdded, arg.CartID, arg.ProductID)
	var i CartDetail
	err := row.Scan(
		&i.ID,
		&i.CartID,
		&i.ProductID,
		&i.QuantityAdded,
	)
	return i, err
}
