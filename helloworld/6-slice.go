package main

import (
	"fmt"
)

func main() {
	var s []int
	fmt.Printf("s=%v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s)) //s=[], len(s)=0, cap(s)=0

	fmt.Printf("s == nil: %v\n", s == nil) // true

	// fmt.Printf("s[0]=%v\n", s[0])          // panic: runtime error: index out of range [0] with length 0

	s1 := make([]int, 3)
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1)) //s1=[   ], len(s1)=3, cap(s1)=3

	fmt.Printf("s1[0]=%v\n", s1[0]) //s1[0]=0

	// fmt.Printf("s1[3]=%v\n", s1[3]) // panic: runtime error: index out of range [3] with length 3

	// s1[3] = 12 //panic: runtime error: index out of range [3] with length 3

	s1 = []int{1, 2, 3, 4, 5}
	fmt.Printf("s1=%v\n", s1) // s1=[1 2 3 4 5]

	s1 = append(s1, 1, 2, 3)
	fmt.Printf("s1=%v\n", s1) // s1=[1 2 3 4 5 1 2 3]

	s2 := []int{0, 1, 2, 3, 4}
	fmt.Println("s2[:3] = ", s2[:3])   // s2[:3] = [0 1 2] 左闭右开
	fmt.Println("s2[3:] = ", s2[3:])   // s2[3:] = [3 4] 左闭右开
	fmt.Println("s2[1:3] = ", s2[1:3]) // s2[1:3] = [1 2] 左闭右开

	var dst1 []int
	copy(dst1, s2)
	fmt.Println("dst1 = ", dst1) // dst1 =  []

	// 注意长度也是类型的一部分（因为slice基于array）

	dst2 := make([]int, len(s2))
	copy(dst2, s2)
	fmt.Println(dst2) // [0 1 2 3 4]

}
