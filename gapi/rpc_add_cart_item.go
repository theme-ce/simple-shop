package gapi

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) AddCartItem(ctx context.Context, req *pb.AddCartItemRequest) (*pb.AddCartItemResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	cart, err := server.store.GetCartByID(ctx, int64(req.GetCartId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %s", err)
	}

	if authPayload.Username != cart.Username {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add item to other user's cart: %s", err)
	}

	arg := db.GetCartDetailParams{
		CartID:    req.GetCartId(),
		ProductID: req.GetProductId(),
	}

	cartDetail, err := server.store.GetCartDetail(ctx, arg)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			createCartDetailArg := db.CreateCartDetailParams{
				CartID:        req.GetCartId(),
				ProductID:     req.GetProductId(),
				QuantityAdded: req.GetQuantityAdded(),
			}

			cartDetail, err = server.store.CreateCartDetail(ctx, createCartDetailArg)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "failed to create cart detail for product id: %d: %s", arg.ProductID, err)
			}
		}
		return nil, status.Errorf(codes.Internal, "failed to get cart detail: %s", err)
	} else {
		updateCartDetailArg := db.UpdateCartDetailParams{
			QuantityAdded: cartDetail.QuantityAdded + req.GetQuantityAdded(),
			CartID:        req.GetCartId(),
			ProductID:     req.GetProductId(),
		}
		cartDetail, err = server.store.UpdateCartDetail(ctx, updateCartDetailArg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update cart detail: %s", err)
		}
	}

	rsp := &pb.AddCartItemResponse{
		Success:       true,
		CartId:        cartDetail.CartID,
		ProductId:     cartDetail.ProductID,
		QuantityAdded: cartDetail.QuantityAdded,
	}
	return rsp, nil
}
