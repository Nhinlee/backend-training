package rpcapi

import (
	"log"

	"habits.com/habit/pb"
)

func (server *Server) Subscribe(req *pb.SubscribeRequest, srv pb.ChatService_SubscribeServer) error {
	log.Printf("start subscribe stream")

	rsp := pb.SubscribeResponse{
		Event: &pb.Event{
			Event: &pb.Event_Message{
				Message: &pb.NewMessageEvent{
					ConversationId: "Conversation ID",
					Content:        "Hello, I'm server",
					MessageId:      "1",
				},
			},
		},
	}

	srv.Send(&rsp)

	return nil
}
