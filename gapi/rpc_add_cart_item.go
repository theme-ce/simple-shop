package gapi

import (
	"context"

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

	arg := db.CreateCartDetailParams{
		CartID:        int64(req.GetCartId()),
		ProductID:     int64(req.GetProductId()),
		QuantityAdded: int64(req.GetQuantityAdded()),
	}

	cartDetail, err := server.store.CreateCartDetail(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create cart detail for product id: %d: %s", arg.ProductID, err)
	}

	rsp := &pb.AddCartItemResponse{
		Success:       true,
		CartId:        int32(cartDetail.CartID),
		ProductId:     int32(cartDetail.ProductID),
		QuantityAdded: int32(cartDetail.QuantityAdded),
	}
	return rsp, nil
}
