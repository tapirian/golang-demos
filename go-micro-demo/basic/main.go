package main

import (
	"context"
	"log"

	"go-micro.dev/v5"
	"go-micro.dev/v5/broker"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Greeting string `json:"greeting"`
}

type Say struct {
	Broker broker.Broker
}

func (s *Say) Hello(ctx context.Context, req *Request, rsp *Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	b := broker.NewHttpBroker(broker.Addrs(":8082"))
	if err := b.Connect(); err != nil {
		log.Fatal(err)
	}
	defer b.Disconnect()

	service := micro.NewService(
		micro.Name("greeter"),
		micro.Address(":8081"),
		micro.Broker(b), // 使用我们自定义的 Broker
	)

	handler := &Say{Broker: b}
	service.Handle(handler)
	service.Run()
}
