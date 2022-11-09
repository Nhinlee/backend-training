package rpcapi

import (
	"context"
	"fmt"
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
			defer wait.Done()
			if conn.active {
				err := conn.stream.Send(&pb.SubscribeResponse{
					Event: event,
				})
				fmt.Printf("Sending message to: %v", conn)

				if err != nil {
					fmt.Printf("Error with stream: %v - error: %v", conn, err)
					conn.active = false
					conn.errors <- err
				}
			}
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
