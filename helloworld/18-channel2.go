package main

import (
	"fmt"
	"time"
)

func main() {
	havecache()
}

// 无缓存的channel
func nocache() {
	ch := make(chan int)
	fmt.Println("start....")
	go func() {
		fmt.Println("write channel...start")
		time.Sleep(1 * time.Second)
		ch <- 1
		ch <- 1
		close(ch)
		fmt.Println("################")
	}()
	fmt.Println("LOCATION 1")
	num := <-ch
	fmt.Println("LOCATION 2")
	fmt.Println("num", num)
	fmt.Println("输出零值： ", <-ch)
	fmt.Println("LOCATION 3")
}

func nocache2() {
	ch := make(chan int)
	go func() {
		fmt.Println("location1")
		ch <- 10
		time.Sleep(1 * time.Second)
		fmt.Println("location2") // 这一行一定会被打印，因为第二个值发送没完成，程序阻塞
		ch <- 20
		time.Sleep(1 * time.Second)
		fmt.Println("location3") // 这一行不会被打印, 接收到通道的第二个值以后，主goroutine就结束了。
		close(ch)
		time.Sleep(1 * time.Second)
		fmt.Println("location4") // 这一行不会被打印，延迟一秒之后。主goroutine已经执行完成了。
	}()
	fmt.Println("waiting for ch...1")
	<-ch
	fmt.Println("waiting for ch...2")
	<-ch
	fmt.Println("waiting for ch...3")
}

// 有缓冲的channel
func havecache() {
	ch := make(chan string, 3)
	go func() {
		ch <- "hello"
		fmt.Println("LOCATION 1")
	}()
	go func() {
		fmt.Println("LOCATION 1")
		ch <- "world"
	}()
	go func() {
		ch <- "hello world"
	}()
	go func() {
		ch <- "china"
	}()
	go func() {
		ch <- "india"
	}()
	go func() {
		ch <- "us"
	}()
	go func() {
		ch <- "japan"
	}()

	fmt.Println(<-ch)
	fmt.Printf("LOC1: cap: %v, len: %v\n", cap(ch), len(ch))
	fmt.Println(<-ch)
	fmt.Printf("LOC2: cap: %v, len: %v\n", cap(ch), len(ch))
	fmt.Println(<-ch)
	fmt.Printf("LOC3: cap: %v, len: %v\n", cap(ch), len(ch))
	fmt.Println(<-ch)
}
