package main

import (
	"fmt"
	"net"
)

func main() {
	// 指定服务器地址和端口
	serverAddr := "127.0.0.1:8088"

	// 创建一个TCP连接
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server")

	// 在这里你可以通过conn进行读写操作，发送和接收数据

	// 示例：发送数据
	message := "Hello, server!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	fmt.Println("Sent:", message)

	// 示例：接收数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error receiving data:", err)
		return
	}
	receivedData := string(buffer[:n])
	fmt.Println("Received:", receivedData)
}
