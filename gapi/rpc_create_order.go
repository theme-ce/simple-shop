package gapi

import (
	"context"

	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot create order on other user's account")
	}

	arg := db.CreateOrderTxParams{
		Username: req.GetUsername(),
	}

	txResult, err := server.store.CreateOrderTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create order: %s", err)
	}

	rsp := &pb.CreateOrderResponse{
		Order: convertOrder(txResult.Order),
	}
	return rsp, nil
}
