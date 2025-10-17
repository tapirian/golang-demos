package lib

import (
	"fmt"
	"sync"
	"time"
)

type DataStore struct {
	mu   sync.RWMutex
	data map[string]int
}

func NewDataStore() *DataStore {
	return &DataStore{
		data: make(map[string]int),
	}
}

// 读操作 - 使用读锁
func (ds *DataStore) Read(id int, key string) {
	fmt.Printf("[读协程 %d] 🔍 尝试获取读锁... (时间: %s)\n", id, time.Now().Format("15:04:05.000"))

	ds.mu.RLock()
	fmt.Printf("[读协程 %d] ✓ 获取读锁成功，开始读取 (时间: %s)\n", id, time.Now().Format("15:04:05.000"))

	// 模拟读取操作
	value := ds.data[key]
	fmt.Printf("[读协程 %d] 读取到 %s = %d\n", id, key, value)
	time.Sleep(1 * time.Second) // 模拟读取耗时

	fmt.Printf("[读协程 %d] 读取完成，释放读锁 (时间: %s)\n", id, time.Now().Format("15:04:05.000"))
	ds.mu.RUnlock()
}

// 写操作 - 使用写锁
func (ds *DataStore) Write(id int, key string, value int) {
	fmt.Printf("[写协程 %d] ✏️  尝试获取写锁... (时间: %s)\n", id, time.Now().Format("15:04:05.000"))

	ds.mu.Lock()
	fmt.Printf("[写协程 %d] ✓ 获取写锁成功，开始写入 (时间: %s)\n", id, time.Now().Format("15:04:05.000"))

	// 模拟写入操作
	ds.data[key] = value
	fmt.Printf("[写协程 %d] 写入 %s = %d\n", id, key, value)
	time.Sleep(1 * time.Second) // 模拟写入耗时

	fmt.Printf("[写协程 %d] 写入完成，释放写锁 (时间: %s)\n", id, time.Now().Format("15:04:05.000"))
	ds.mu.Unlock()
}

func RWMutexCall() {
	fmt.Println("========== RWMutex 读写锁演示 ==========\n")

	store := NewDataStore()
	var wg sync.WaitGroup

	// 先初始化一些数据
	store.data["counter"] = 0

	fmt.Println("【场景1】多个读操作可以并发执行")
	fmt.Println("启动 3 个读协程...\n")

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			store.Read(id, "counter")
		}(i)
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n--------------------------------------------------")
	fmt.Println("【场景2】写操作会阻塞所有读和写")
	fmt.Println("启动 1 个写协程和 2 个读协程...\n")

	// 先启动写操作
	wg.Add(1)
	go func() {
		defer wg.Done()
		store.Write(1, "counter", 100)
	}()

	time.Sleep(200 * time.Millisecond)

	// 再启动读操作（会被写锁阻塞）
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			store.Read(id, "counter")
		}(i)
		time.Sleep(100 * time.Millisecond)
	}

	wg.Wait()
	time.Sleep(500 * time.Millisecond)

	fmt.Println("\n--------------------------------------------------")
	fmt.Println("【场景3】读操作期间，写操作必须等待")
	fmt.Println("启动 2 个读协程和 1 个写协程...\n")

	// 先启动读操作
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			store.Read(id, "counter")
		}(i)
		time.Sleep(100 * time.Millisecond)
	}

	time.Sleep(200 * time.Millisecond)

	// 再启动写操作（会被读锁阻塞）
	wg.Add(1)
	go func() {
		defer wg.Done()
		store.Write(2, "counter", 200)
	}()

	wg.Wait()

	fmt.Printf("\n========== 总结 ==========\n")
	fmt.Println("✓ 多个读锁可以同时持有（并发读取）")
	fmt.Println("✓ 写锁是独占的，会阻塞所有读和写")
	fmt.Println("✓ 读锁存在时，写操作必须等待")
	fmt.Println("✓ 适用于读多写少的场景")
	fmt.Printf("\n最终数据: counter = %d\n", store.data["counter"])
}
