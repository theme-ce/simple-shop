package gapi

import (
	"context"
	"errors"

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
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "cart is not existed: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to get cart: %s", err)
	}

	if authPayload.Username != cart.Username {
		return nil, status.Errorf(codes.PermissionDenied, "cannot add item to other user's cart: %s", err)
	}

	arg := db.AddCartItemTxParams{
		CartId:    req.GetCartId(),
		ProductId: req.GetProductId(),
		Quantity:  req.GetQuantityAdded(),
	}

	cartDetail, err := server.store.AddCartItemTx(ctx, arg)
	if err != nil {
		return nil, err
	}

	rsp := &pb.AddCartItemResponse{
		Success:       true,
		CartId:        cartDetail.CartId,
		ProductId:     cartDetail.ProductId,
		QuantityAdded: cartDetail.Quantity,
	}
	return rsp, nil
}
