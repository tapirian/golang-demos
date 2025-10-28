package main

import (
	"log"

	"go-micro.dev/v5/broker"
)

func main() {
	b := broker.NewHttpBroker(broker.Addrs("127.0.0.1:8082"))
	if err := b.Connect(); err != nil {
		log.Fatal(err)
	}
	defer b.Disconnect()

	// 订阅主题
	println("sub2开始订阅主题 greeter.topic")
	_, err := b.Subscribe("greeter.topic", func(p broker.Event) error {
		msg := p.Message()
		log.Printf("收到消息: %s\n", string(msg.Body))
		return nil
	})
	if err != nil {
		log.Fatalf("订阅失败: %v", err)
	}

	<-make(chan struct{})
}
