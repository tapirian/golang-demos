package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for scanner.Scan() {
			fmt.Printf("server说：%s\n", scanner.Text())
		}
	}()

	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
