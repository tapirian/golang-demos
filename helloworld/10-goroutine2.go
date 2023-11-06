package main

func main() {
	go golist()
	forlist(100)
}

func golist() {
	for i := 0; i < 10000000; i++ {
		println("delay process: ", i)
	}
}

func forlist(x int) {
	for i := 0; i < x; i++ {
		println("main process: ", i)
	}
}
