package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("\n=== 无缓冲通道示例3: 工作协调 ===")
	ch := make(chan int)

	// 工作者
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("工作者: 处理任务 %d\n", i)
			time.Sleep(500 * time.Millisecond)
			ch <- i // 发送结果，等待确认
			fmt.Printf("工作者: 任务 %d 已确认\n", i)
		}
		close(ch)
	}()

	// 协调者
	for result := range ch {
		fmt.Printf("协调者: 收到结果 %d，进行验证\n", result)
		time.Sleep(300 * time.Millisecond)
	}
}
