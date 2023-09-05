package gapi

import (
	"context"
	"database/sql"
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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestLoginUser(t *testing.T) {
	user, password := randomUser(t)
	accessPayloadId, err := uuid.NewRandom()
	accessToken := "accessToken"
	require.NoError(t, err)
	accessPayload := &token.Payload{
		ID:        accessPayloadId,
		Username:  user.Username,
		IsAdmin:   false,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Minute),
	}
	refreshPayloadId, err := uuid.NewRandom()
	refreshToken := "refreshToken"
	require.NoError(t, err)
	refreshPayload := &token.Payload{
		ID:        refreshPayloadId,
		Username:  user.Username,
		IsAdmin:   false,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Minute * 15),
	}
	session := db.Session{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
		CreatedAt:    refreshPayload.IssuedAt,
	}

	testCases := []struct {
		name          string
		req           *pb.LoginUserRequest
		buildStubs    func(store *mockdb.MockStore, tokenMaker *mocktoken.MockMaker)
		checkResponse func(t *testing.T, res *pb.LoginUserResponse, err error)
	}{
		{
			name: "OK",
			req: &pb.LoginUserRequest{
				Username: user.Username,
				Password: password,
			},
			buildStubs: func(store *mockdb.MockStore, tokenMaker *mocktoken.MockMaker) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(user, nil)
				tokenMaker.EXPECT().
					CreateToken(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(accessToken, accessPayload, nil)
				tokenMaker.EXPECT().
					CreateToken(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(1).
					Return(refreshToken, refreshPayload, nil)
				store.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(1).
					Return(session, nil)
			},
			checkResponse: func(t *testing.T, res *pb.LoginUserResponse, err error) {
				require.NoError(t, err)
				require.NotNil(t, res)
				createdUser := res.GetUser()
				require.Equal(t, user.Username, createdUser.Username)
				require.Equal(t, accessToken, res.AccessToken)
				require.Equal(t, refreshToken, res.RefreshToken)
			},
		},
		{
			name: "InternalError",
			req: &pb.LoginUserRequest{
				Username: user.Username,
				Password: password,
			},
			buildStubs: func(store *mockdb.MockStore, tokenMaker *mocktoken.MockMaker) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
				tokenMaker.EXPECT().
					CreateToken(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
				tokenMaker.EXPECT().
					CreateToken(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *pb.LoginUserResponse, err error) {
				require.Error(t, err)
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, codes.Internal, st.Code())
			},
		},
		{
			name: "UserNotFound",
			req: &pb.LoginUserRequest{
				Username: user.Username,
				Password: password,
			},
			buildStubs: func(store *mockdb.MockStore, tokenMaker *mocktoken.MockMaker) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.Username)).
					Times(1).
					Return(db.User{}, db.ErrRecordNotFound)
				tokenMaker.EXPECT().
					CreateToken(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
				tokenMaker.EXPECT().
					CreateToken(gomock.Any(), gomock.Any(), gomock.Any()).
					Times(0)
				store.EXPECT().
					CreateSession(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, res *pb.LoginUserResponse, err error) {
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
			tokenMaker := mocktoken.NewMockMaker(ctrl)

			tc.buildStubs(store, tokenMaker)
			server := NewTestServerWithTokenMaker(t, store, tokenMaker)

			res, err := server.LoginUser(context.Background(), tc.req)
			tc.checkResponse(t, res, err)
		})
	}
}
