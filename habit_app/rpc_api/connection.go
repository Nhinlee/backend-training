package rpcapi

import "habits.com/habit/pb"

type Connection struct {
	stream pb.ChatService_SubscribeServer
	id     int64
	active bool
	errors chan error
}
