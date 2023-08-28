package gapi

import (
	db "github.com/theme-ce/simple-shop/db/sqlc"
	"github.com/theme-ce/simple-shop/pb"
	"github.com/theme-ce/simple-shop/util"
)

type Server struct {
	pb.UnimplementedSimpleShopServer
	config util.Config
	store  db.Store
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
