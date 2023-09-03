package gapi

import (
	"context"

	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateOrderStatus(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateOrderRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if !authPayload.IsAdmin {
		return nil, status.Errorf(codes.PermissionDenied, "only admin allow to update order status")
	}

	arg := db.UpdateOrderStatusParams{
		Status: req.GetStatus(),
		ID:     req.GetId(),
	}

	order, err := server.store.UpdateOrderStatus(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update order status: %s", err)
	}

	rsp := &pb.UpdateOrderResponse{
		Order: convertOrder(order),
	}
	return rsp, nil
}

func validateUpdateOrderRequest(req *pb.UpdateOrderRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateOrderStatus(req.Status); err != nil {
		violations = append(violations, fieldViolation("order_status", err))
	}

	return violations
}
