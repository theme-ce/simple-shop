package db

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddCartItemTxParams struct {
	CartId    int64
	ProductId int64
	Quantity  int64
}

type AddCartItemTxResponse struct {
	Success   bool
	CartId    int64
	ProductId int64
	Quantity  int64
}

func (store *SQLStore) AddCartItemTx(ctx context.Context, req AddCartItemTxParams) (AddCartItemTxResponse, error) {
	var result AddCartItemTxResponse
	err := store.execTx(ctx, func(q *Queries) error {
		arg := GetCartDetailParams{
			CartID:    req.CartId,
			ProductID: req.ProductId,
		}

		cartDetail, err := store.GetCartDetail(ctx, arg)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				createCartDetailArg := CreateCartDetailParams{
					CartID:        req.CartId,
					ProductID:     req.ProductId,
					QuantityAdded: req.Quantity,
				}

				cartDetail, err = store.CreateCartDetail(ctx, createCartDetailArg)
				if err != nil {
					return status.Errorf(codes.Internal, "failed to create cart detail for product id: %d: %s", arg.ProductID, err)
				}
			}
			return status.Errorf(codes.Internal, "failed to get cart detail: %s", err)
		} else {
			updateCartDetailArg := UpdateCartDetailParams{
				QuantityAdded: cartDetail.QuantityAdded + req.Quantity,
				CartID:        req.CartId,
				ProductID:     req.ProductId,
			}
			cartDetail, err = store.UpdateCartDetail(ctx, updateCartDetailArg)
			if err != nil {
				return status.Errorf(codes.Internal, "failed to update cart detail: %s", err)
			}
		}

		result = AddCartItemTxResponse{
			CartId:    cartDetail.CartID,
			ProductId: cartDetail.ProductID,
			Quantity:  cartDetail.QuantityAdded,
		}

		return nil
	})

	return result, err
}
