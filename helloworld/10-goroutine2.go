package main

import "time"

func main() {
	done := make(chan bool, 1) // 带缓冲的通道

	go func() {
		golist()
		done <- true
	}()

	forlist(100) // 立即执行，不会阻塞
	<-done       // 最后检查协程是否完成
}

func golist() {
	for i := 0; i < 100; i++ {
		println("delay process: ", i)
		time.Sleep(time.Microsecond * 100)
	}
}

func forlist(x int) {
	for i := 0; i < x; i++ {
		println("main process: ", i)
		time.Sleep(time.Microsecond * 100)
	}
}
