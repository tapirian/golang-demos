package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	const n = 46
	fibN := fib(n) // 比较慢，等待执行
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// 开启协程包含死循环会导致该协程一直运行，直到程序中止。（不会影响主程序运行）
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
