package gapi

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	mockdb "github.com/theme-ce/simple-shop/db/mock"
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/token"
)

func TestUpdateItem(t *testing.T) {
	username := "testtest"

	cartDetail := db.CartDetail{
		ID:            1,
		CartID:        1,
		ProductID:     1,
		QuantityAdded: 2,
	}
	newQuantity := int64(5)

	testCases := []struct {
		name          string
		req           *pb.UpdateCartItemQuantityRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.UpdateCartItemQuantityResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.UpdateCartItemQuantityRequest{
				CartId:      cartDetail.CartID,
				ProductId:   cartDetail.ProductID,
				NewQuantity: newQuantity,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateCartDetailParams{
					CartID:        cartDetail.CartID,
					ProductID:     cartDetail.ProductID,
					QuantityAdded: newQuantity,
				}

				store.EXPECT().
					GetCartByID(gomock.Any(), gomock.Eq(cartDetail.CartID)).
					Times(1).
					Return(db.Cart{
						ID:       cartDetail.CartID,
						Username: username,
					}, nil)
				store.EXPECT().
					UpdateCartDetail(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.CartDetail{
						ID:            1,
						CartID:        cartDetail.CartID,
						ProductID:     cartDetail.ProductID,
						QuantityAdded: newQuantity,
					}, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateCartItemQuantityResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, true, res.Success)
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
			res, err := server.UpdateCartItemQuantity(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
