package gapi

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	mockdb "github.com/theme-ce/simple-shop/db/mock"
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/token"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestUpdateProduct(t *testing.T) {
	username := "testtest"
	product := db.Product{
		ID:            1,
		Name:          "testtest",
		Description:   "testtest",
		Price:         1000,
		StockQuantity: 100,
	}

	newName := "new_product_name"

	testCases := []struct {
		name          string
		req           *pb.UpdateProductRequest
		buildStubs    func(store *mockdb.MockStore)
		buildContext  func(t *testing.T, tokenMaker token.Maker) context.Context
		checkResponse func(t *testing.T, res *pb.UpdateProductResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.UpdateProductRequest{
				Id:   product.ID,
				Name: &newName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateProductParams{
					ID: product.ID,
					Name: pgtype.Text{
						String: newName,
						Valid:  true,
					},
				}

				store.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Product{
						ID:            product.ID,
						Name:          newName,
						Description:   product.Description,
						Price:         product.Price,
						StockQuantity: product.StockQuantity,
					}, nil)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateProductResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.NotEqual(t, product.Name, res.Product.Name)
				require.Equal(t, product.Description, res.Product.Description)
				require.Equal(t, product.Price, res.Product.Price)
				require.Equal(t, product.StockQuantity, res.Product.StockQuantity)
			},
		},
		{
			name: "InternalError",
			req: &pb.UpdateProductRequest{
				Id:   product.ID,
				Name: &newName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateProductParams{
					ID: product.ID,
					Name: pgtype.Text{
						String: newName,
						Valid:  true,
					},
				}

				store.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(arg)).
					Times(1).
					Return(db.Product{}, sql.ErrConnDone)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateProductResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "Unauthenticated",
			req: &pb.UpdateProductRequest{
				Id:   product.ID,
				Name: &newName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateProductParams{
					ID: product.ID,
					Name: pgtype.Text{
						String: newName,
						Valid:  true,
					},
				}

				store.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, true, -time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateProductResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Unauthenticated, st.Code())
			},
		},
		{
			name: "PermissionDenied",
			req: &pb.UpdateProductRequest{
				Id:   product.ID,
				Name: &newName,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.UpdateProductParams{
					ID: product.ID,
					Name: pgtype.Text{
						String: newName,
						Valid:  true,
					},
				}

				store.EXPECT().
					UpdateProduct(gomock.Any(), gomock.Eq(arg)).
					Times(0)
			},
			buildContext: func(t *testing.T, tokenMaker token.Maker) context.Context {
				return newContextWithBearerToken(t, tokenMaker, username, false, time.Minute)
			},
			checkResponse: func(t *testing.T, res *pb.UpdateProductResponse, err error) {
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
			res, err := server.UpdateProduct(ctx, tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
