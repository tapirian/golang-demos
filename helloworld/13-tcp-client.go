package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("client on: %v\n", conn.LocalAddr())
	fmt.Printf("server on: %v\n", conn.RemoteAddr())
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
