package gapi

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/theme-ce/simple-shop/db/mock"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteProduct(t *testing.T) {
	username := "testtest"
	productId := int64(1)

	testCases := []struct {
		name          string
		req           *pb.DeleteProductRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.DeleteProductResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.DeleteProductRequest{
				Id: productId,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Eq(productId)).
					Times(1).
					Return(nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteProductResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, true, res.Success)
				require.Equal(t, fmt.Sprintf("product id %d is successfully delete", productId), res.Message)
			},
		},
		{
			name: "Unauthenticated",
			req: &pb.DeleteProductRequest{
				Id: productId,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteProductResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "PermissionDenied",
			req: &pb.DeleteProductRequest{
				Id: productId,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteProductResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.PermissionDenied, st.Code())
			},
		},
		{
			name: "InternalError",
			req: &pb.DeleteProductRequest{
				Id: productId,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteProduct(gomock.Any(), gomock.Any()).
					Times(1).
					Return(sql.ErrConnDone)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteProductResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
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
			res, err := server.DeleteProduct(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
