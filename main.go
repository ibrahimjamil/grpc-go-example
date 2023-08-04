package main

import (
	"context"
	"grpc_example/invoicer"
	"net"

	"github.com/pastelnetwork/gonode/common/log"
	"google.golang.org/grpc"
)

type myInvoiceServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoiceServer) Create(context.Context, *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {
	return &invoicer.CreateResponse{
		Pdf:  []byte("Hello"),
		Docs: []byte("Hello"),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatal(err)
	}
	serverRegistrar := grpc.NewServer()

	service := &myInvoiceServer{}
	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
