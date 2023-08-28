package gapi

import (
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		Address:           user.Address.String,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
		IsAdmin:           user.IsAdmin,
	}
}
