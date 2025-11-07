package lib

import (
	"fmt"
	"sync"
	"time"
)

func SyncCondCall() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	dataReady := false

	// 消费者（等待条件）
	go func() {
		mu.Lock()
		for !dataReady { // 条件不满足就等待
			fmt.Println("Consumer: waiting for data...")
			cond.Wait() // 等待时会自动释放锁
		}
		fmt.Println("Consumer: data is ready!")
		mu.Unlock()
	}()

	// 生产者（修改条件并通知）
	time.Sleep(1 * time.Second)
	mu.Lock()
	fmt.Println("Producer: preparing data...")
	dataReady = true
	mu.Unlock()

	cond.Signal() // 通知一个等待的 goroutine

	time.Sleep(500 * time.Millisecond)
	fmt.Println("All done")
}
