package main

import "fmt"

func main() {
	var ch chan int
	if ch == nil {
		fmt.Println("channel is null")
	} else {
		fmt.Printf("channel: %v", ch)
	}
	ch = make(chan int)

	ch1 := ch
	go func() {
		fmt.Print("hello")
		ch <- 1
		close(ch)
	}()
	if ch1 == ch {
		fmt.Println("eq")
	} else {
		fmt.Println("neq")
	}
	fmt.Print("world")
	println("out", <-ch)
	println("out", <-ch1)

	go func() {
	}()
	fmt.Print("!")

	if ch == nil {
		fmt.Println("null")
	}
}
