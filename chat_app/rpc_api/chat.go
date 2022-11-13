package rpcapi

import (
	"context"
	"fmt"
	"sync"
	"v1/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	messageEvent := &pb.NewMessageEvent{
		Content:        req.Content,
		MessageId:      "MessageID",
		ConversationId: req.ConversationId,
	}
	event := &pb.Event{Event: &pb.Event_Message{Message: messageEvent}}

	err := server.sendMessageToConversation(ctx, req.ConversationId, event)
	if err != nil {
		return nil, err
	}

	resp := &pb.SendMessageResponse{
		MessageId: "MessageID",
	}
	return resp, nil
}

func (server *Server) sendMessageToConversation(ctx context.Context, conversationID string, message *pb.Event) error {
	wait := sync.WaitGroup{}
	done := make(chan int)

	// Get all user id by conversation id
	userIDs, err := server.store.ListUserIdByConversationId(ctx, conversationID)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to get user ids by conversation id")
	}
	fmt.Printf("user ids: %v\n", userIDs)

	// Send message to all user have active connection in conversation
	for _, userID := range userIDs {

		conn := server.connections[userID]
		if conn == nil {
			continue
		}

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
		}(message, conn)
	}

	go func() {
		wait.Wait()
		close(done)
	}()
	<-done

	return nil
}
