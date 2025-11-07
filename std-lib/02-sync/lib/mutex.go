package lib

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) SlowIncrement(id int) {
	fmt.Printf("[协程 %d] 尝试获取锁... (时间: %s)\n", id, time.Now().Format("15:04:05.000"))

	c.mu.Lock()
	fmt.Printf("[协程 %d] ✓ 成功获取锁，开始工作 (时间: %s)\n", id, time.Now().Format("15:04:05.000"))

	// 模拟耗时操作
	c.value++
	fmt.Printf("[协程 %d] 正在处理数据... 当前值=%d\n", id, c.value)
	time.Sleep(2 * time.Second) // 持有锁 2 秒

	fmt.Printf("[协程 %d] 工作完成，释放锁 (时间: %s)\n", id, time.Now().Format("15:04:05.000"))
	c.mu.Unlock()
}

func MutexCall() {
	fmt.Println("========== Mutex 阻塞效果演示 ==========")
	fmt.Println()

	counter := &Counter{}
	var wg sync.WaitGroup

	// 启动 5 个协程同时竞争锁
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			counter.SlowIncrement(id)
		}(i)

		// 稍微错开启动时间，让输出更清晰
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()

	fmt.Printf("\n========== 完成 ==========\n")
	fmt.Printf("最终计数值: %d\n", counter.value)
	fmt.Println("\n观察要点:")
	fmt.Println("1. 只有一个协程能获取锁并执行")
	fmt.Println("2. 其他协程必须等待，直到锁被释放")
	fmt.Println("3. 每个协程持有锁约 2 秒，其他协程被阻塞")
	fmt.Println("4. 总执行时间约为 10 秒 (5个协程 × 2秒)")
}
