package rpcapi

import (
	"context"
	"sync"
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

func (server *Server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	// Filter all connections with user_id in conversation

	event := &pb.NewMessageEvent{
		Content:        req.Content,
		MessageId:      "MessageID",
		ConversationId: req.ConversationId,
	}

	wait := sync.WaitGroup{}
	done := make(chan int)

	for _, conn := range server.connections {
		wait.Add(1)

		go func(event *pb.Event, conn *Connection) {
			// TODO: implement send message logic here
		}(&pb.Event{Event: &pb.Event_Message{Message: event}}, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()
	<-done

	resp := &pb.SendMessageResponse{
		MessageId: "MessageID",
	}
	return resp, nil
}
