package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a:", a) //a: [0 0 0 0 0]

	a[1] = 123
	fmt.Println("a:", a) //a: [0 123 0 0 0]

	fmt.Println("len(a):", len(a)) // len(a): 5

	b := [5]int{1, 2, 3, 5}
	fmt.Println("b:", b) //b: [1 2 3 5 0]

	var bb [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			bb[i][j] = i + j
		}
	}

	fmt.Println(bb) // [[0 1 2] [1 2 3]]
}
