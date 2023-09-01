package gapi

import (
	"context"
	"fmt"

	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateCartItemQuantity(ctx context.Context, req *pb.UpdateCartItemQuantityRequest) (*pb.UpdateCartItemQuantityResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	cart, err := server.store.GetCartByID(ctx, int64(req.GetCartId()))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %s", err)
	}

	if authPayload.Username != cart.Username {
		return nil, status.Errorf(codes.PermissionDenied, "cannot update item in other user's cart: %s", err)
	}

	arg := db.UpdateCartDetailParams{
		QuantityAdded: req.GetNewQuantity(),
		CartID:        req.GetCartId(),
		ProductID:     req.GetProductId(),
	}

	cartDetail, err := server.store.UpdateCartDetail(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update cart detail: %s", err)
	}

	rsp := &pb.UpdateCartItemQuantityResponse{
		Success: true,
		Message: fmt.Sprintf("successfully update product id: %d in cart id: %d", cartDetail.ProductID, cartDetail.CartID),
	}
	return rsp, nil
}
