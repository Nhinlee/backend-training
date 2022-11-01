package rpcapi

import (
	"context"
	"fmt"
	"sync"

	"habits.com/habit/pb"
)

func (server *Server) Subscribe(req *pb.SubscribeRequest, srv pb.ChatService_SubscribeServer) error {
	conn := &Connection{
		stream: srv,
		id:     req.UserID,
		active: true,
		errors: make(chan error),
	}

	server.connections = append(server.connections, conn)

	// go func() {
	// 	for i := 1; i < 10; i++ {
	// 		log.Printf("LOOPING")
	// 		time.Sleep(time.Second * 1)
	// 		srv.Send(&pb.SubscribeResponse{
	// 			Event: &pb.Event{
	// 				Event: &pb.Event_Message{
	// 					Message: &pb.NewMessageEvent{
	// 						ConversationId: "Conversation ID",
	// 						Content:        "Hello, I'm server",
	// 						MessageId:      fmt.Sprintf("%d", i),
	// 					},
	// 				},
	// 			},
	// 		})
	// 	}
	// }()

	return <-conn.errors
}

func (server *Server) BroadcastMessage(context context.Context, event *pb.Event) (*pb.Close, error) {
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
					fmt.Printf("Error with stream: %v - Error: %v ", conn, err)
					conn.active = false
					conn.errors <- err
				}
			}
		}(event, conn)

	}

	go func() {
		wait.Wait()
		close(done)
	}()

	<-done
	return &pb.Close{}, nil
}
