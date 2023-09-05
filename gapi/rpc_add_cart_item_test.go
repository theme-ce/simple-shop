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

func TestAddCartItem(t *testing.T) {
	username := "testtest"
	cartId := int64(1)
	productId := int64(1)
	quantity := int64(3)

	testCases := []struct {
		name          string
		req           *pb.AddCartItemRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.AddCartItemResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.AddCartItemRequest{
				CartId:        cartId,
				ProductId:     productId,
				QuantityAdded: quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.AddCartItemTxParams{
					CartId:    cartId,
					ProductId: productId,
					Quantity:  quantity,
				}

				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Eq(cartId)).
					Times(1).
					Return(db.Cart{
						ID:       cartId,
						Username: username,
					}, nil)
				store.EXPECT().
					AddCartItemTx(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.AddCartItemTxResponse{
						Success:   true,
						CartId:    cartId,
						ProductId: productId,
						Quantity:  quantity,
					}, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.AddCartItemResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, cartId, res.CartId)
				require.Equal(t, productId, res.ProductId)
				require.Equal(t, quantity, res.QuantityAdded)
			},
		},
		{
			name: "NoAuthenticated",
			req: &pb.AddCartItemRequest{
				CartId:        cartId,
				ProductId:     productId,
				QuantityAdded: quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					AddCartItemTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.AddCartItemResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &pb.AddCartItemRequest{
				CartId:        cartId,
				ProductId:     productId,
				QuantityAdded: quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Cart{}, sql.ErrConnDone)
				store.EXPECT().
					AddCartItemTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.AddCartItemResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "CartNotExisted",
			req: &pb.AddCartItemRequest{
				CartId:        cartId,
				ProductId:     productId,
				QuantityAdded: quantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Any()).
					Times(1).
					Return(db.Cart{}, db.ErrRecordNotFound)
				store.EXPECT().
					AddCartItemTx(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.AddCartItemResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.NotFound, st.Code())
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
			res, err := server.AddCartItem(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
