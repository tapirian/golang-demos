package main

import "fmt"

func main() {
	natural := make(chan int)
	square := make(chan int)
	go counter(natural)
	go squares(natural, square)
	print(square)
}

func counter1() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

func counter2() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	func() {
		for {
			x, ok := <-squares
			if !ok {
				break
			}
			fmt.Println(x)
		}
	}()
}

// 单向channel
func counter(natural chan<- int) {
	for x := 0; x < 100; x++ {
		natural <- x
	}
	close(natural)
}

func squares(natural <-chan int, square chan<- int) {
	for x := range natural {
		square <- x * x
	}
	close(square)
}

func print(square <-chan int) {
	for x := range square {
		fmt.Println(x)
	}
}
