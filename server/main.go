package main

import (
	"log"
	"net"

	pb "grpc_example/proto"

	"google.golang.org/grpc"
)

type Server struct {
	pb.GreetServiceServer
}

func main() {
	port := ":8080"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &Server{})
	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}
