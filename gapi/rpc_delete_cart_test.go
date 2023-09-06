package gapi

import (
	"context"
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

func TestDeleteCart(t *testing.T) {
	username := "testtest"
	invalidUsername := "1"

	testCases := []struct {
		name          string
		req           *pb.DeleteCartRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.DeleteCartResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.DeleteCartRequest{
				Username: username,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteCart(gomock.Any(), gomock.Eq(username)).
					Times(1).
					Return(nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteCartResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, true, res.Success)
			},
		},
		{
			name: "InvalidUsername",
			req: &pb.DeleteCartRequest{
				Username: invalidUsername,
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					DeleteCart(gomock.Any(), gomock.Any()).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, invalidUsername, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.DeleteCartResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.InvalidArgument, st.Code())
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
			res, err := server.DeleteCart(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
