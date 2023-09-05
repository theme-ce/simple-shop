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

func TestCreateCart(t *testing.T) {
	username := "testtest"
	cartId := int64(1)

	testCases := []struct {
		name          string
		req           *pb.CreateCartRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.CreateCartResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.CreateCartRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCart(gomock.Any(), gomock.Eq(username)).
					Times(1).
					Return(db.Cart{ID: int64(cartId)}, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.CreateCartResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, cartId, res.CartId)
			},
		},
		{
			name: "PermissionDenied",
			req: &pb.CreateCartRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCart(gomock.Any(), gomock.Eq(username)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, "permission_denied", false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.CreateCartResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.PermissionDenied, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &pb.CreateCartRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCart(gomock.Any(), gomock.Eq(username)).
					Times(1).
					Return(db.Cart{}, sql.ErrConnDone)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.CreateCartResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "NoAuthenticated",
			req: &pb.CreateCartRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					CreateCart(gomock.Any(), gomock.Eq(username)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.CreateCartResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
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
			res, err := server.CreateCart(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
