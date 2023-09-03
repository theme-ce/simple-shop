package gapi

import (
	"context"

	"github.com/theme-ce/simple-shop/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetCartItems(ctx context.Context, req *pb.GetCartItemsRequest) (*pb.GetCartItemsResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(err)
	}

	if authPayload.Username != req.GetUsername() {
		return nil, status.Errorf(codes.PermissionDenied, "cannot get items from other user's cart: %s", err)
	}

	cart, err := server.store.GetCartByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart: %s", err)
	}

	cartDetails, err := server.store.GetCartDetailsByCartID(ctx, cart.ID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get cart details: %s", err)
	}

	cartItems := []*pb.CartItem{}

	for _, detail := range cartDetails {
		product, err := server.store.GetProduct(ctx, detail.ProductID)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get product: %s", err)
		}

		cartItem := &pb.CartItem{
			ProductId:          detail.ProductID,
			ProductName:        product.Name,
			ProductDescription: product.Description,
			Price:              product.Price,
			Quantity:           detail.QuantityAdded,
		}

		cartItems = append(cartItems, cartItem)
	}

	rsp := &pb.GetCartItemsResponse{
		CartItems: cartItems,
	}
	return rsp, nil
}
