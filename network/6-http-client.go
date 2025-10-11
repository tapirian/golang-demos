package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	var data []byte
	data = []byte("server,你好，我是client")
	req, err := http.NewRequest(http.MethodPost, "http://127.0.0.1:9000", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	req.Header.Set("content-type", "application/text")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	buffer := make([]byte, 4096)
	n, err := res.Body.Read(buffer)
	defer res.Body.Close()
	if err != io.EOF {
		panic(err)
	}

	fmt.Println("响应： ", string(buffer[:n]))

	fmt.Println("end......")
}
