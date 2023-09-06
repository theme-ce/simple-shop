package gapi

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/theme-ce/simple-shop/db/mock"
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRemoveCartItem(t *testing.T) {
	username := "testtest"
	cartId := int64(1)
	productId := int64(1)

	testCases := []struct {
		name          string
		req           *pb.RemoveCartItemRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.RemoveCartItemResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.RemoveCartItemRequest{
				CartId:    cartId,
				ProductId: productId,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.DeleteCartDetailParams{
					CartID:    cartId,
					ProductID: productId,
				}

				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Eq(cartId)).
					Times(1).
					Return(db.Cart{
						ID:       cartId,
						Username: username,
					}, nil)
				store.EXPECT().
					DeleteCartDetail(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.RemoveCartItemResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, true, res.Success)
			},
		},
		{
			name: "InternalError",
			req: &pb.RemoveCartItemRequest{
				CartId:    cartId,
				ProductId: productId,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Cart{}, sql.ErrConnDone)
				store.EXPECT().
					DeleteCartDetail(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.RemoveCartItemResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "Unauthenticated",
			req: &pb.RemoveCartItemRequest{
				CartId:    cartId,
				ProductId: productId,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					DeleteCartDetail(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.RemoveCartItemResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "PermissionDenied",
			req: &pb.RemoveCartItemRequest{
				CartId:    cartId,
				ProductId: productId,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Eq(cartId)).
					Times(1).
					Return(db.Cart{
						ID:       cartId,
						Username: username,
					}, nil)
				store.EXPECT().
					DeleteCartDetail(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, "invalid-username", false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.RemoveCartItemResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.PermissionDenied, st.Code())
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)

			tc.buildStubs(store)
			server := newTestServer(t, store)

			ctx := tc.buildContext(t, server.tokenMaker)
			res, err := server.RemoveCartItem(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
