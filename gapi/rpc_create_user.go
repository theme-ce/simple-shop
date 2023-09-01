package gapi

import (
	"context"

	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/util"
	"github.com/theme-ce/simple-shop/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:       req.GetUsername(),
			HashedPassword: hashedPassword,
			FirstName:      req.GetFirstName(),
			LastName:       req.GetLastName(),
			Email:          req.GetEmail(),
		},
	}

	txResult, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %s", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(txResult.User),
	}

	return resp, nil
}

func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	if err := val.ValidateFirstName(req.GetFirstName()); err != nil {
		violations = append(violations, fieldViolation("first_name", err))
	}
	if err := val.ValidateLastName(req.GetLastName()); err != nil {
		violations = append(violations, fieldViolation("last_name", err))
	}
	if err := val.ValidateAdress(req.GetAddress()); err != nil {
		violations = append(violations, fieldViolation("address", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	return violations
}
