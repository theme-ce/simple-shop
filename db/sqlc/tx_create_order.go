package db

import (
	"context"
)

type CreateOrderTxParams struct {
	Username string
}

type CreateOrderTxResult struct {
	Order Order
}

func (store *SQLStore) CreateOrderTx(ctx context.Context, arg CreateOrderTxParams) (CreateOrderTxResult, error) {
	var result CreateOrderTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		cart, err := store.GetCartByUsername(ctx, arg.Username)
		if err != nil {
			return err
		}

		cartDetails, err := store.GetCartDetailsByCartID(ctx, cart.ID)
		if err != nil {
			return err
		}

		totalPrice := 0.0
		for _, item := range cartDetails {
			product, err := store.GetProduct(ctx, item.ProductID)
			if err != nil {
				return err
			}

			totalPrice += float64(product.Price) * float64(item.QuantityAdded)
		}

		arg := CreateOrderParams{
			Username:   arg.Username,
			TotalPrice: totalPrice,
			Status:     "PENDING",
		}

		result.Order, err = store.CreateOrder(ctx, arg)
		if err != nil {
			return err
		}

		for _, item := range cartDetails {
			product, err := store.GetProduct(ctx, item.ProductID)
			if err != nil {
				return err
			}

			_, err = store.CreateOrderDetail(ctx, CreateOrderDetailParams{
				ProductID:          item.ProductID,
				Username:           arg.Username,
				QuantityOrdered:    item.QuantityAdded,
				PriceAtTimeOfOrder: product.Price,
			})
			if err != nil {
				return err
			}
		}

		return store.ClearCart(ctx, cart.ID)
	})

	return result, err
}
