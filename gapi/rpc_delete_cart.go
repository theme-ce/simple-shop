package gapi

import (
	"context"

	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteCart(ctx context.Context, req *pb.DeleteCartRequest) (*pb.DeleteCartResponse, error) {
	violations := validateDeleteCartRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot delete other user's cart: %s", err)
	}

	err = server.store.DeleteCart(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete cart: %s", err)
	}

	rsp := &pb.DeleteCartResponse{
		Success: true,
	}
	return rsp, nil
}

func validateDeleteCartRequest(req *pb.DeleteCartRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	return violations
}
