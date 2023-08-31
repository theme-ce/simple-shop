package gapi

import (
	"context"
	"fmt"

	"github.com/theme-ce/simple-shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if !authPayload.IsAdmin {
		return nil, status.Errorf(codes.PermissionDenied, "only admin allow to delete product")
	}

	product_id := req.GetId()
	err = server.store.DeleteProduct(ctx, product_id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete product: %s", err)
	}

	rsp := &pb.DeleteProductResponse{
		Success: true,
		Message: fmt.Sprintf("product id %d is successfully delete", product_id),
	}

	return rsp, nil
}
