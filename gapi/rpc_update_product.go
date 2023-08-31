package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateUpdateProductRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if !authPayload.IsAdmin {
		return nil, status.Errorf(codes.PermissionDenied, "only admin allow to create product")
	}

	arg := db.UpdateProductParams{
		Name: pgtype.Text{
			String: req.GetName(),
			Valid:  req.Name != nil,
		},
		Description: pgtype.Text{
			String: req.GetDescription(),
			Valid:  req.Description != nil,
		},
		Price: pgtype.Int8{
			Int64: req.GetPrice(),
			Valid: req.Price != nil,
		},
		StockQuantity: pgtype.Int8{
			Int64: req.GetStockQuantity(),
			Valid: req.StockQuantity != nil,
		},
		ID: req.GetId(),
	}

	product, err := server.store.UpdateProduct(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update product: %s", err)
	}

	rsp := &pb.UpdateProductResponse{
		Product: convertProduct(product),
	}
	return rsp, nil
}

func validateUpdateProductRequest(req *pb.UpdateProductRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateProductName(req.GetName()); err != nil {
		violations = append(violations, fieldViolation("name", err))
	}
	if err := val.ValidateProductPrice(int(req.GetPrice())); err != nil {
		violations = append(violations, fieldViolation("price", err))
	}
	if err := val.ValidateProductQuantity(int(req.GetStockQuantity())); err != nil {
		violations = append(violations, fieldViolation("product_quantity", err))
	}
	return violations
}
