package gapi

import (
	"context"

	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := server.store.CreateUser(ctx, db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: req.Password,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		Email:          req.Email,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create user: %s", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return resp, nil
}
