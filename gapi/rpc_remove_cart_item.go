package gapi

import (
	"context"

	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) RemoveCartItem(ctx context.Context, req *pb.RemoveCartItemRequest) (*pb.RemoveCartItemResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	cart, err := server.store.GetCartByID(ctx, req.GetCartId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %s", err)
	}

	if authPayload.Username != cart.Username {
		return nil, status.Errorf(codes.PermissionDenied, "cannot remove item from other user's cart: %s", err)
	}

	arg := db.DeleteCartDetailParams{
		CartID:    req.GetCartId(),
		ProductID: req.GetProductId(),
	}

	err = server.store.DeleteCartDetail(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to remove cart item: %s", err)
	}

	rsp := &pb.RemoveCartItemResponse{
		Success: true,
	}
	return rsp, nil
}
