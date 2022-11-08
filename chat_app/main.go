package main

import (
	"log"
	"net"
	"v1/config"
	"v1/pb"
	rpcapi "v1/rpc_api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configure: ", err)
	}

	runGRPCServer(&config)
}

func runGRPCServer(config *config.Config) {
	server, err := rpcapi.NewServer(config)
	if err != nil {
		log.Fatal("Cannot create server: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

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
