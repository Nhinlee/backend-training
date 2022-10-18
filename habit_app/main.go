package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"habits.com/habit/api"
	db "habits.com/habit/db/sqlc"
	"habits.com/habit/pb"
	rpcapi "habits.com/habit/rpc_api"
	"habits.com/habit/utils"
)

func main() {
	// Load configs
	config, err := utils.LoadConfig(".")
	if err != nil {
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	runGrpcServer(&config, store)
}

func runGrpcServer(config *utils.Config, store db.Store) {
	server, err := rpcapi.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserModifierServer(grpcServer, server)
	pb.RegisterChatServiceServer(grpcServer, server)
	reflection.Register(grpcServer) // Share api visible for client - consider to security or not?

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server")
	}
}

func runHTTPServer(config *utils.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
