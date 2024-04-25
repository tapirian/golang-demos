package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	testselect()
}

func countdown() {
	fmt.Println("count down")

	tick := time.Tick(1 * time.Second)
	for count := 10; count > 0; count-- {
		fmt.Println(count)
		<-tick
	}

	// for {
	// 	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	// 	<-tick
	// }
}

// 当有标准输入的时候，程序退出
func countdown2() {
	fmt.Println("count down2")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
}

func testselect() {
	// 创建两个通道
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动两个 goroutine 分别向通道发送数据
	go func() {
		for {
			ch1 <- "from ch1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			ch2 <- "from ch2"
			time.Sleep(time.Second * 3)
		}
	}()

	// 使用 select 语句处理多个通道的消息
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received:", msg2)
		}
	}
}
