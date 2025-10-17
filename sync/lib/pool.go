package lib

import (
	"fmt"
	"sync"
)

func SyncPollCall() {
	// 创建对象池
	pool := sync.Pool{
		New: func() any {
			fmt.Println("Creating new object")
			buf := make([]byte, 1024) // 1KB 临时缓冲
			return &buf
		},
	}

	// 从池中获取对象
	obj1 := pool.Get().(*[]byte)
	fmt.Println("Got object 1:", len(*obj1))

	// 使用完放回池
	pool.Put(obj1)

	// 再次获取对象（会复用上次的 obj1）
	obj2 := pool.Get().(*[]byte)
	fmt.Println("Got object 2:", len(*obj2))

	// 获取新的对象（池空时会调用 New）
	pool.Put(nil)
	obj3 := pool.Get().(*[]byte)
	fmt.Println("Got object 3:", len(*obj3))
}
