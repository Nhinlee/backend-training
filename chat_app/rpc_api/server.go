package rpcapi

import "v1/config"

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) (*Server, error) {
	server := &Server{
		config: config,
	}

	return server, nil
}
