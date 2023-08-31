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

func (server *Server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	violations := validateCreateProductRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	if !authPayload.IsAdmin {
		return nil, status.Errorf(codes.PermissionDenied, "only admin allow to create product")
	}

	arg := db.CreateProductParams{
		Name:          req.GetName(),
		Description:   req.GetDescription(),
		Price:         int64(req.GetPrice()),
		StockQuantity: int64(req.GetStockQuantity()),
	}

	product, err := server.store.CreateProduct(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot create product: %s", err)
	}

	rsp := &pb.CreateProductResponse{
		Product: convertProduct(product),
	}
	return rsp, nil
}

func validateCreateProductRequest(req *pb.CreateProductRequest) (violations []*errdetails.BadRequest_FieldViolation) {
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
