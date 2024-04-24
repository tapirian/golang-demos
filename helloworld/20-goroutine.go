package main

import (
	"fmt"
	"sync"
)

func main() {
	// goroutineleak()
	goroutineWaitGroup()
}

// 模拟一种goroutine泄露，
// 这是一个没有缓冲的channel,
// 当程序返回时，没有goroutine排空channel剩余值，再往这个channel写入时，会永远阻塞下去
// 解决办法： 将channle变成一个有缓冲的channel
func goroutineleak() {
	var country []string
	country = append(country, "China", "India", "Japan", "US", "UK")

	// country := []string{"China", "India", "Japan", "US", "UK"}

	ch := make(chan string)
	for _, c := range country {
		go func(c string) {
			ch <- c
			fmt.Println(c)
		}(c)
	}

	for range country {
		if c := <-ch; c != "China" {
			fmt.Println("结束了：", c)
			return
		}
	}
}

// sync.WaitGroup
func goroutineWaitGroup() {
	var country []string
	country = append(country, "China", "India", "Japan", "US", "UK")

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, c := range country {
		wg.Add(1)
		go func(cou string) {
			defer wg.Done()
			// 其他操作
			ch <- cou
		}(c)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}
}
