package rpcapi

import (
	"v1/config"
	db "v1/db/sqlc"
	"v1/pb"
)

type Server struct {
	// Platform
	config *config.Config
	store  db.Store

	// gRPC
	pb.UnimplementedChatServiceServer

	// Chat
	connections map[string]*Connection
}

func NewServer(config *config.Config, store db.Store) (*Server, error) {

	connections := make(map[string]*Connection)
	server := &Server{
		config:      config,
		connections: connections,
		store:       store,
	}

	return server, nil
}
