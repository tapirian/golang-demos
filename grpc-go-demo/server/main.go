package main

import (
	"context"
	"fmt"
	hello "grpc-go-demo/proto"
	"net"

	"google.golang.org/grpc"
)

func main() {
	grpcAddr := "127.0.0.1:40001"
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	grpcServer := grpc.NewServer()
	hello.RegisterHelloGRPCServer(grpcServer, &HelloServer{})
	if err := grpcServer.Serve(listener); err != nil {
		listener.Close()
		panic(fmt.Sprintf("failed to serve: %v", err))
	}
}

type HelloServer struct {
}

func (s *HelloServer) SayHello(ctx context.Context, req *hello.SayHelloRequest) (*hello.SayHelloReply, error) {
	return &hello.SayHelloReply{
		Content: "Hello " + req.Name,
	}, nil
}
