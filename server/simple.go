package main

import (
	"context"

	pb "grpc_example/proto"
)

func (s *Server) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
