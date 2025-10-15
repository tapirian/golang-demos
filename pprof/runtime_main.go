package main

import (
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// ===== CPU Profiling =====
	f, _ := os.Create("cpu.prof")
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	go MockCPUUse()
	go MockMemoryUse()
	time.Sleep(5 * time.Second)

	// ===== Memory Profiling =====
	f2, _ := os.Create("mem.prof")
	pprof.WriteHeapProfile(f2)
	f2.Close()
}

func MockCPUUse() {
	for {
		for i := 0; i < 1000000; i++ {
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// 模拟内存使用持续增加
func MockMemoryUse() {
	var data [][]byte
	for {
		row := make([]byte, 10*1024) // byte固定1个字节，10*1024也就是10kB
		data = append(data, row)
		log.Printf("total size = %dKB\n", len(data)*10)
		time.Sleep(100 * time.Millisecond)
	}
}
