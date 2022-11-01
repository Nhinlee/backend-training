package rpcapi

import (
	"fmt"

	db "habits.com/habit/db/sqlc"
	"habits.com/habit/pb"
	"habits.com/habit/token"
	"habits.com/habit/utils"
)

type Server struct {
	pb.UnimplementedUserModifierServer
	pb.UnimplementedChatServiceServer
	config       *utils.Config
	store        db.Store
	tokenFactory token.TokenFactory

	// Chat
	connections []*Connection
}

func NewServer(config *utils.Config, store db.Store) (*Server, error) {
	tokenFactory, err := token.NewPasetoTokenFactory(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token factory: %w", err)
	}

	connections := make([]*Connection, 0)

	server := &Server{
		store:        store,
		tokenFactory: tokenFactory,
		config:       config,
		connections:  connections,
	}

	return server, nil
}
