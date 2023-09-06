package gapi

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	mockdb "github.com/theme-ce/simple-shop/db/mock"
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/token"
	mocktoken "github.com/theme-ce/simple-shop/token/mock"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestRenewAccessToken(t *testing.T) {
	username := "testtest"
	refreshToken := "refresh-token"
	refreshPayload := &token.Payload{
		ID:        uuid.New(),
		Username:  username,
		IsAdmin:   false,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 15),
	}
	accessToken := "access-token"
	accessPayload := &token.Payload{
		ID:        uuid.New(),
		Username:  username,
		IsAdmin:   false,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Minute),
	}

	testCases := []struct {
		name          string
		req           *pb.RenewAccessTokenRequest
		buildStubs    func(store *mockdb.MockStore, tokenMaker *mocktoken.MockMaker)
		checkResponse func(t *testing.T, res *pb.RenewAccessTokenResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.RenewAccessTokenRequest{
				RefreshToken: refreshToken,
			},
			buildStubs: func(store *mockdb.MockStore, tokenMaker *mocktoken.MockMaker) {
				tokenMaker.EXPECT().
					VerifyToken(gomock.Eq(refreshToken)).
					Times(1).
					Return(refreshPayload, nil)
				store.EXPECT().
					GetSession(gomock.Any(), gomock.Eq(refreshPayload.ID)).
					Times(1).
					Return(db.Session{
						ID:           refreshPayload.ID,
						Username:     refreshPayload.Username,
						RefreshToken: refreshToken,
						UserAgent:    "",
						ClientIp:     "",
						IsBlocked:    false,
						ExpiredAt:    refreshPayload.ExpiredAt,
					}, nil)
				tokenMaker.EXPECT().
					CreateToken(gomock.Eq(refreshPayload.Username), gomock.Eq(refreshPayload.IsAdmin), time.Minute).
					Times(1).
					Return(accessToken, accessPayload, nil)
			},
			checkResponse: func(t *testing.T, res *pb.RenewAccessTokenResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				require.Equal(t, accessToken, res.AccessToken)
				require.Equal(t, timestamppb.New(accessPayload.ExpiredAt), res.AccessTokenExpiredAt)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockStore(ctrl)
			tokenMaker := mocktoken.NewMockMaker(ctrl)

			tc.buildStubs(store, tokenMaker)
			server := NewTestServerWithTokenMaker(t, store, tokenMaker)

			res, err := server.RenewAccessToken(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
