package rpcapi

import (
	"v1/config"
	"v1/pb"
)

type Server struct {
	config *config.Config
	pb.UnimplementedChatServiceServer
	connections map[string]*Connection
}

func NewServer(config *config.Config) (*Server, error) {

	connections := make(map[string]*Connection)
	server := &Server{
		config:      config,
		connections: connections,
	}

	return server, nil
}
