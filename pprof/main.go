package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	go func() {
		var mux = http.NewServeMux()
		// http.ListenAndServe("localhost:6060", nil)
		// mux.HandleFunc("/debug/pprof/", http.DefaultServeMux.ServeHTTP)
		mux.Handle("/debug/pprof/", http.DefaultServeMux)
		http.ListenAndServe("localhost:6060", mux)
	}()
	println("服务已运行")
	// go MockMemoryUse()
	// go MockCPUUse()
	// go MockGoroutineLeak()
	go MockGCfreq()

	for {
		time.Sleep(2 * time.Second)
		fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	}
}

// 模拟频繁GC
func MockGCfreq() {
	for {
		_ = make([]byte, 10*1024*1024)
		time.Sleep(100 * time.Millisecond)
	}
}

// 模拟协程泄露
func GoroutineLeak(ch <-chan struct{}) {
	for range ch { // ch 被 close 时 loop 结束
		// 如果确实需要并发处理可在此派发短生命周期的 goroutine
		go func() {
			// ...短任务
		}()
		time.Sleep(1 * time.Second)
	}
}
func MockGoroutineLeak() {
	for {
		ch := make(chan struct{})
		go GoroutineLeak(ch)
		time.Sleep(500 * time.Millisecond)
	}
}

// 模拟内存使用持续增加
func MockMemoryUse() {
	var data [][]byte
	for {
		row := make([]byte, 10*1024*1024)
		data = append(data, row)
		// log.Printf("total size = %dKB\n", len(data)*10)
		time.Sleep(100 * time.Microsecond)
	}
}

// 模拟CPU使用率持续增加
func MockCPUUse() {
	for {
		for i := 0; i < 10000000; i++ {
		}
		time.Sleep(1 * time.Second)
	}
}
