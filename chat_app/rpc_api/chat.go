package rpcapi

import (
	"v1/pb"
)

func (server *Server) Subscribe(req *pb.SubscribeRequest, srv pb.ChatService_SubscribeServer) error {
	conn := &Connection{
		stream: srv,
		active: true,
		errors: make(chan error),
	}

	server.connections[req.UserId] = conn

	return <-conn.errors
}
