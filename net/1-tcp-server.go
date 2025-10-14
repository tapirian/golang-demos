package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// 在这里你可以通过conn进行读写操作，处理客户端的请求

	// 示例：接收客户端发送的数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	receivedData := string(buffer[:n])
	fmt.Println("Received from client:", receivedData)

	// 示例：向客户端发送数据
	message := "Hello, client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
	fmt.Println("Sent to client:", message)
}

func main() {
	// 指定服务器监听的地址和端口
	listenAddr := "0.0.0.0:8088"

	// 创建一个TCP监听器
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on", listenAddr)

	// 无限循环等待客户端连接
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Client connected:", conn.RemoteAddr())

		// 启动一个新的 goroutine 处理客户端连接
		go handleConnection(conn)
	}
}
