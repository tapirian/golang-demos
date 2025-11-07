package main

import (
	"context"
	"fmt"
	"log"
	"time"

	hello "grpc-go-demo-client/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:40001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := hello.NewHelloGRPCClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.SayHello(ctx, &hello.SayHelloRequest{
		Name: "World",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	fmt.Printf("Response: %s\n", resp.Content)
}
