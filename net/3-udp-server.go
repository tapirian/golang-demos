package main

import (
	"log"
	"net"
)

func main() {
	addr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	println("udp server on: ", conn.LocalAddr().String())
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println(err)
		}
		println("接收到：", string(buffer[:n]))
	}
}
