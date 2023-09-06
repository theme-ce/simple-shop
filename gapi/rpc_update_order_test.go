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
	"github.com/theme-ce/simple-shop/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUpdateOrder(t *testing.T) {
	username := "testtest"

	order := db.Order{
		ID:         1,
		Username:   username,
		TotalPrice: 2000,
		Status:     util.Pending,
	}

	newStatus := util.Shipping

	testCases := []struct {
		name          string
		req           *pb.UpdateOrderRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.UpdateOrderResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.UpdateOrderRequest{
				Id:     order.ID,
				Status: newStatus,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateOrderStatusParams{
					ID:     order.ID,
					Status: newStatus,
				}

				store.EXPECT().
					UpdateOrderStatus(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Order{
						ID:         order.ID,
						Username:   username,
						TotalPrice: order.TotalPrice,
						Status:     newStatus,
					}, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateOrderResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEqual(t, order.Status, res.Order.Status)
				require.Equal(t, order.ID, res.Order.Id)
				require.Equal(t, order.Username, res.Order.Username)
				require.Equal(t, order.TotalPrice, res.Order.TotalPrice)
			},
		},
		{
			name: "InternalError",
			req: &pb.UpdateOrderRequest{
				Id:     order.ID,
				Status: newStatus,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateOrderStatusParams{
					ID:     order.ID,
					Status: newStatus,
				}

				store.EXPECT().
					UpdateOrderStatus(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Order{}, sql.ErrConnDone)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "Unauthenticated",
			req: &pb.UpdateOrderRequest{
				Id:     order.ID,
				Status: newStatus,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateOrderStatusParams{
					ID:     order.ID,
					Status: newStatus,
				}

				store.EXPECT().
					UpdateOrderStatus(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateOrderResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "PermissionDenied",
			req: &pb.UpdateOrderRequest{
				Id:     order.ID,
				Status: newStatus,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateOrderStatusParams{
					ID:     order.ID,
					Status: newStatus,
				}

				store.EXPECT().
					UpdateOrderStatus(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateOrderResponse, err error) {
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
			res, err := server.UpdateOrderStatus(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
