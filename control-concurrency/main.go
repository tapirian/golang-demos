package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	total := 20
	batchSize := 10
	fmt.Println(total, batchSize)
	RateControl(total, batchSize)
}

func batchChuckReq(total int, batchSize int) {
	for i := 0; i < total; i += batchSize {
		end := i + batchSize
		if end > total {
			end = total
		}

		var wg sync.WaitGroup
		for j := i; j < end; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				exec(j)
			}(j)
		}
		wg.Wait() // 等这一批结束
		time.Sleep(time.Second * 1)
	}
}

func bufferChannelReq(total int, batchSize int) {
	ch := make(chan int, batchSize)
	var wg sync.WaitGroup
	for i := 0; i < total; i++ {
		wg.Add(1)
		// 占用一个位置
		fmt.Println("i=", i)
		ch <- i
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			// 每次调用释放一个位置
			fmt.Printf("i=%d, ch=%d\n", i, <-ch)
			exec(i)
		}(i)
	}
	wg.Wait()
}
func exec(j int) {
	// fmt.Println(time.Now().Format(time.DateTime), j)
	fmt.Println(time.Now(), j)
}

func RateControl(total int, rateNum int) {
	limiter := rate.NewLimiter(5, 10) // 每秒10个请求，桶容量20
	for i := 0; i < 30; i++ {
		limiter.Wait(context.Background()) // 等待许可
		exec(i)
	}
}
