package main

import (
	"encoding/json"
	"log"
	"time"

	"go-micro.dev/v5/broker"
)

type Message struct {
	Name string `json:"name"`
}

func PubMsg(b broker.Broker, topic string, msg any) (err error) {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("消息序列化失败%v: \n", err)
		return err
	}

	err = b.Publish(topic, &broker.Message{
		Header: map[string]string{"Content-Type": "application/json"},
		Body:   data,
	})
	if err != nil {
		log.Printf("发布消息失败%v: \n", err)
		return err
	}
	println("发布消息: ", string(data))
	return
}

func main() {
	b := broker.NewHttpBroker(broker.Addrs(":8082"))
	if err := b.Connect(); err != nil {
		log.Fatal(err)
	}
	defer b.Disconnect()

	for {
		msg := Message{Name: "World" + time.Now().Format("15:04:05")}
		PubMsg(b, "greeter.topic", msg)
		time.Sleep(5 * time.Second)
	}
}
