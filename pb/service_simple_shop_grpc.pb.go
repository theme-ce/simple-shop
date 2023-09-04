// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: service_simple_shop.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SimpleShop_CreateUser_FullMethodName             = "/pb.SimpleShop/CreateUser"
	SimpleShop_LoginUser_FullMethodName              = "/pb.SimpleShop/LoginUser"
	SimpleShop_UpdateUser_FullMethodName             = "/pb.SimpleShop/UpdateUser"
	SimpleShop_CreateProduct_FullMethodName          = "/pb.SimpleShop/CreateProduct"
	SimpleShop_UpdateProduct_FullMethodName          = "/pb.SimpleShop/UpdateProduct"
	SimpleShop_DeleteProduct_FullMethodName          = "/pb.SimpleShop/DeleteProduct"
	SimpleShop_CreateCart_FullMethodName             = "/pb.SimpleShop/CreateCart"
	SimpleShop_GetCartItems_FullMethodName           = "/pb.SimpleShop/GetCartItems"
	SimpleShop_DeleteCart_FullMethodName             = "/pb.SimpleShop/DeleteCart"
	SimpleShop_AddCartItem_FullMethodName            = "/pb.SimpleShop/AddCartItem"
	SimpleShop_UpdateCartItemQuantity_FullMethodName = "/pb.SimpleShop/UpdateCartItemQuantity"
	SimpleShop_RemoveCartItem_FullMethodName         = "/pb.SimpleShop/RemoveCartItem"
	SimpleShop_CreateOrder_FullMethodName            = "/pb.SimpleShop/CreateOrder"
	SimpleShop_UpdateOrderStatus_FullMethodName      = "/pb.SimpleShop/UpdateOrderStatus"
	SimpleShop_RenewAccessToken_FullMethodName       = "/pb.SimpleShop/RenewAccessToken"
)

// SimpleShopClient is the client API for SimpleShop service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleShopClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error)
	DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error)
	CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CreateCartResponse, error)
	GetCartItems(ctx context.Context, in *GetCartItemsRequest, opts ...grpc.CallOption) (*GetCartItemsResponse, error)
	DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*DeleteCartResponse, error)
	AddCartItem(ctx context.Context, in *AddCartItemRequest, opts ...grpc.CallOption) (*AddCartItemResponse, error)
	UpdateCartItemQuantity(ctx context.Context, in *UpdateCartItemQuantityRequest, opts ...grpc.CallOption) (*UpdateCartItemQuantityResponse, error)
	RemoveCartItem(ctx context.Context, in *RemoveCartItemRequest, opts ...grpc.CallOption) (*RemoveCartItemResponse, error)
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	UpdateOrderStatus(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error)
}

type simpleShopClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleShopClient(cc grpc.ClientConnInterface) SimpleShopClient {
	return &simpleShopClient{cc}
}

func (c *simpleShopClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, SimpleShop_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, SimpleShop_LoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, SimpleShop_UpdateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, SimpleShop_CreateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*UpdateProductResponse, error) {
	out := new(UpdateProductResponse)
	err := c.cc.Invoke(ctx, SimpleShop_UpdateProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) DeleteProduct(ctx context.Context, in *DeleteProductRequest, opts ...grpc.CallOption) (*DeleteProductResponse, error) {
	out := new(DeleteProductResponse)
	err := c.cc.Invoke(ctx, SimpleShop_DeleteProduct_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) CreateCart(ctx context.Context, in *CreateCartRequest, opts ...grpc.CallOption) (*CreateCartResponse, error) {
	out := new(CreateCartResponse)
	err := c.cc.Invoke(ctx, SimpleShop_CreateCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) GetCartItems(ctx context.Context, in *GetCartItemsRequest, opts ...grpc.CallOption) (*GetCartItemsResponse, error) {
	out := new(GetCartItemsResponse)
	err := c.cc.Invoke(ctx, SimpleShop_GetCartItems_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) DeleteCart(ctx context.Context, in *DeleteCartRequest, opts ...grpc.CallOption) (*DeleteCartResponse, error) {
	out := new(DeleteCartResponse)
	err := c.cc.Invoke(ctx, SimpleShop_DeleteCart_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) AddCartItem(ctx context.Context, in *AddCartItemRequest, opts ...grpc.CallOption) (*AddCartItemResponse, error) {
	out := new(AddCartItemResponse)
	err := c.cc.Invoke(ctx, SimpleShop_AddCartItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) UpdateCartItemQuantity(ctx context.Context, in *UpdateCartItemQuantityRequest, opts ...grpc.CallOption) (*UpdateCartItemQuantityResponse, error) {
	out := new(UpdateCartItemQuantityResponse)
	err := c.cc.Invoke(ctx, SimpleShop_UpdateCartItemQuantity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) RemoveCartItem(ctx context.Context, in *RemoveCartItemRequest, opts ...grpc.CallOption) (*RemoveCartItemResponse, error) {
	out := new(RemoveCartItemResponse)
	err := c.cc.Invoke(ctx, SimpleShop_RemoveCartItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, SimpleShop_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) UpdateOrderStatus(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, SimpleShop_UpdateOrderStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleShopClient) RenewAccessToken(ctx context.Context, in *RenewAccessTokenRequest, opts ...grpc.CallOption) (*RenewAccessTokenResponse, error) {
	out := new(RenewAccessTokenResponse)
	err := c.cc.Invoke(ctx, SimpleShop_RenewAccessToken_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleShopServer is the server API for SimpleShop service.
// All implementations must embed UnimplementedSimpleShopServer
// for forward compatibility
type SimpleShopServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error)
	DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error)
	CreateCart(context.Context, *CreateCartRequest) (*CreateCartResponse, error)
	GetCartItems(context.Context, *GetCartItemsRequest) (*GetCartItemsResponse, error)
	DeleteCart(context.Context, *DeleteCartRequest) (*DeleteCartResponse, error)
	AddCartItem(context.Context, *AddCartItemRequest) (*AddCartItemResponse, error)
	UpdateCartItemQuantity(context.Context, *UpdateCartItemQuantityRequest) (*UpdateCartItemQuantityResponse, error)
	RemoveCartItem(context.Context, *RemoveCartItemRequest) (*RemoveCartItemResponse, error)
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	UpdateOrderStatus(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error)
	mustEmbedUnimplementedSimpleShopServer()
}

// UnimplementedSimpleShopServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleShopServer struct {
}

func (UnimplementedSimpleShopServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedSimpleShopServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedSimpleShopServer) UpdateUser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedSimpleShopServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedSimpleShopServer) UpdateProduct(context.Context, *UpdateProductRequest) (*UpdateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedSimpleShopServer) DeleteProduct(context.Context, *DeleteProductRequest) (*DeleteProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProduct not implemented")
}
func (UnimplementedSimpleShopServer) CreateCart(context.Context, *CreateCartRequest) (*CreateCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedSimpleShopServer) GetCartItems(context.Context, *GetCartItemsRequest) (*GetCartItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartItems not implemented")
}
func (UnimplementedSimpleShopServer) DeleteCart(context.Context, *DeleteCartRequest) (*DeleteCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (UnimplementedSimpleShopServer) AddCartItem(context.Context, *AddCartItemRequest) (*AddCartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCartItem not implemented")
}
func (UnimplementedSimpleShopServer) UpdateCartItemQuantity(context.Context, *UpdateCartItemQuantityRequest) (*UpdateCartItemQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartItemQuantity not implemented")
}
func (UnimplementedSimpleShopServer) RemoveCartItem(context.Context, *RemoveCartItemRequest) (*RemoveCartItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCartItem not implemented")
}
func (UnimplementedSimpleShopServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedSimpleShopServer) UpdateOrderStatus(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderStatus not implemented")
}
func (UnimplementedSimpleShopServer) RenewAccessToken(context.Context, *RenewAccessTokenRequest) (*RenewAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RenewAccessToken not implemented")
}
func (UnimplementedSimpleShopServer) mustEmbedUnimplementedSimpleShopServer() {}

// UnsafeSimpleShopServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleShopServer will
// result in compilation errors.
type UnsafeSimpleShopServer interface {
	mustEmbedUnimplementedSimpleShopServer()
}

func RegisterSimpleShopServer(s grpc.ServiceRegistrar, srv SimpleShopServer) {
	s.RegisterService(&SimpleShop_ServiceDesc, srv)
}

func _SimpleShop_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_LoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_UpdateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_CreateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_UpdateProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).UpdateProduct(ctx, req.(*UpdateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_DeleteProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).DeleteProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_DeleteProduct_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).DeleteProduct(ctx, req.(*DeleteProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_CreateCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).CreateCart(ctx, req.(*CreateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_GetCartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).GetCartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_GetCartItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).GetCartItems(ctx, req.(*GetCartItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_DeleteCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).DeleteCart(ctx, req.(*DeleteCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_AddCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).AddCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_AddCartItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).AddCartItem(ctx, req.(*AddCartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_UpdateCartItemQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartItemQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).UpdateCartItemQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_UpdateCartItemQuantity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).UpdateCartItemQuantity(ctx, req.(*UpdateCartItemQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_RemoveCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCartItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).RemoveCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_RemoveCartItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).RemoveCartItem(ctx, req.(*RemoveCartItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_UpdateOrderStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).UpdateOrderStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_UpdateOrderStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).UpdateOrderStatus(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleShop_RenewAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RenewAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleShopServer).RenewAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SimpleShop_RenewAccessToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleShopServer).RenewAccessToken(ctx, req.(*RenewAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SimpleShop_ServiceDesc is the grpc.ServiceDesc for SimpleShop service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SimpleShop_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.SimpleShop",
	HandlerType: (*SimpleShopServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _SimpleShop_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _SimpleShop_LoginUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _SimpleShop_UpdateUser_Handler,
		},
		{
			MethodName: "CreateProduct",
			Handler:    _SimpleShop_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _SimpleShop_UpdateProduct_Handler,
		},
		{
			MethodName: "DeleteProduct",
			Handler:    _SimpleShop_DeleteProduct_Handler,
		},
		{
			MethodName: "CreateCart",
			Handler:    _SimpleShop_CreateCart_Handler,
		},
		{
			MethodName: "GetCartItems",
			Handler:    _SimpleShop_GetCartItems_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _SimpleShop_DeleteCart_Handler,
		},
		{
			MethodName: "AddCartItem",
			Handler:    _SimpleShop_AddCartItem_Handler,
		},
		{
			MethodName: "UpdateCartItemQuantity",
			Handler:    _SimpleShop_UpdateCartItemQuantity_Handler,
		},
		{
			MethodName: "RemoveCartItem",
			Handler:    _SimpleShop_RemoveCartItem_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _SimpleShop_CreateOrder_Handler,
		},
		{
			MethodName: "UpdateOrderStatus",
			Handler:    _SimpleShop_UpdateOrderStatus_Handler,
		},
		{
			MethodName: "RenewAccessToken",
			Handler:    _SimpleShop_RenewAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_simple_shop.proto",
}
