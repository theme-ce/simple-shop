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

func TestGetCartItems(t *testing.T) {
	username := "testtest"
	cart := db.Cart{
		ID:       1,
		Username: username,
	}
	cartDetails := []db.CartDetail{
		{
			ID:            1,
			CartID:        1,
			ProductID:     1,
			QuantityAdded: 2,
		},
		{
			ID:            2,
			CartID:        1,
			ProductID:     2,
			QuantityAdded: 3,
		},
	}
	product1 := db.Product{
		ID:            1,
		Name:          "test1",
		Description:   "test1",
		Price:         1000,
		StockQuantity: 100,
	}
	product2 := db.Product{
		ID:            2,
		Name:          "test2",
		Description:   "test2",
		Price:         2000,
		StockQuantity: 200,
	}

	testCases := []struct {
		name          string
		req           *pb.GetCartItemsRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.GetCartItemsResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.GetCartItemsRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByUsername(gomock.Any(), gomock.Eq(username)).
					Times(1).
					Return(cart, nil)
				store.EXPECT().
					GetCartDetailsByCartID(gomock.Any(), gomock.Eq(cart.ID)).
					Times(1).
					Return(cartDetails, nil)
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product1.ID)).
					Times(1).
					Return(product1, nil)
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product2.ID)).
					Times(1).
					Return(product2, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetCartItemsResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Len(t, res.CartItems, 2)
				require.Equal(t, product1.ID, res.CartItems[0].ProductId)
				require.Equal(t, product2.ID, res.CartItems[1].ProductId)
			},
		},
		{
			name: "InternalError",
			req: &pb.GetCartItemsRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByUsername(gomock.Any(), gomock.Eq(username)).
					Times(1).
					Return(db.Cart{}, sql.ErrConnDone)
				store.EXPECT().
					GetCartDetailsByCartID(gomock.Any(), gomock.Eq(cart.ID)).
					Times(0)
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product1.ID)).
					Times(0)
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Eq(product2.ID)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetCartItemsResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "Unauthenticated",
			req: &pb.GetCartItemsRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByUsername(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					GetCartDetailsByCartID(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetCartItemsResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "PermissionDenied",
			req: &pb.GetCartItemsRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetCartByUsername(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					GetCartDetailsByCartID(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					GetProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, "invalid-username", false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.GetCartItemsResponse, err error) {
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
			res, err := server.GetCartItems(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
