package rpcapi

import "v1/pb"

type Connection struct {
	stream pb.ChatService_SubscribeServer
	active bool
	errors chan error
}
