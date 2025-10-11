package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handConn(conn)
	}
}

func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
}

func handConn(conn net.Conn) {
	defer conn.Close()

	// var wg sync.WaitGroup
	// wg.Add(1)
	go func() {
		// defer wg.Done()
		input := bufio.NewScanner(conn)
		for input.Scan() {
			// go echo(conn, input.Text()+"\n", 1*time.Second)
			fmt.Fprintln(os.Stdout, "client说：", input.Text())
		}
	}()
	io.Copy(conn, os.Stdin)
	// go func() {
	// 	defer wg.Done()
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	for scanner.Scan() {
	// 		// fmt.Fprintln(os.Stdout, "server说：", scanner.Text())
	// 		fmt.Fprintln(conn, scanner.Text())
	// 	}
	// }()
	// wg.Wait()
}
