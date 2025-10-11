package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n=== 有缓冲通道示例3: 批量处理 ===")
	batchSize := 5
	ch := make(chan int, batchSize)
	done := make(chan bool)

	// 数据收集器
	go func() {
		batch := make([]int, 0, batchSize)
		for num := range ch {
			batch = append(batch, num)
			if len(batch) == batchSize {
				fmt.Printf("处理批次: %v\n", batch)
				batch = batch[:0] // 清空批次
			}
		}
		if len(batch) > 0 {
			fmt.Printf("处理最后批次: %v\n", batch)
		}
		done <- true
	}()

	// 发送数据
	for i := 1; i <= 12; i++ {
		ch <- i
	}
	close(ch)
	<-done
}
