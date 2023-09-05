package gapi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/token"
	"github.com/theme-ce/simple-shop/util"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func NewTestServerWithTokenMaker(t *testing.T, store db.Store, tokenMaker token.Maker) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	server.tokenMaker = tokenMaker
	require.NoError(t, err)

	return server
}
