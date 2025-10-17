package lib

import "sync"

func OnceCall() {
	onceBody := func() {
		println("OnceBody Called")
	}

	var once sync.Once
	done := make(chan struct{})
	for i := 0; i < 10; i++ {
		go func(i int) {
			println("goroutine called: ", i)
			once.Do(onceBody)
			println("goroutine ended: ", i)
			done <- struct{}{}
		}(i)
	}

	for i := 0; i < 10; i++ {
		println("main goroutine waiting: ", i)
		<-done
	}
}
