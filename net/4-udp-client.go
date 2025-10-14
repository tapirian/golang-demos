package main

import (
	"log"
	"net"
)

func main() {
	raddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	laddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8081,
	}

	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	println("udp client on: ", conn.LocalAddr().String())
	messge := "server你好，我是client"
	_, err = conn.Write([]byte(messge))
	if err != nil {
		log.Println(err)
	}
}
