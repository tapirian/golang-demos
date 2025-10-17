package lib

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCall() {
	var counter int32 = 0
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1) // 原子自增
			}
			fmt.Printf("Goroutine %d done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("Final counter:", atomic.LoadInt32(&counter))
}
