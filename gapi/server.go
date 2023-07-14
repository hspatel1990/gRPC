package gapi

import (
	db "github.com/hspatel1990/gRPC/db/sqlc"
	"github.com/hspatel1990/gRPC/pb"
	"github.com/hspatel1990/gRPC/util"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config util.Config
	store  db.Store
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store) (*Server, error) {

	server := &Server{
		config: config,
		store:  store,
	}

	return server, nil
}
