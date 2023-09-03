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

func convertProduct(product db.Product) *pb.Product {
	return &pb.Product{
		Name:          product.Name,
		Description:   product.Description,
		Price:         product.Price,
		StockQuantity: product.StockQuantity,
		CreatedAt:     timestamppb.New(product.CreatedAt),
	}
}

func convertOrder(order db.Order) *pb.Order {
	return &pb.Order{
		Id:         order.ID,
		Username:   order.Username,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
	}
}
