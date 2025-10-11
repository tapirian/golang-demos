package main

import (
	"fmt"
)

func main() {
	// new 初始化
	ar := *new([]int)
	fmt.Printf("ar=%v, len(ar)=%v, cap(ar)=%v\n", ar, len(ar), cap(ar)) //ar=[], len(ar)=0, cap(ar)=0
	fmt.Printf("ar == nil: %v\n", ar == nil)                            // false
	ar = []int{1234, 12345}
	fmt.Printf("ar=%v\n", ar)

	// make初始化
	arr := make([]int, 4, 5)
	// arr[3] = 1 // panic: runtime error: index out of range [3] with length 3
	// arr=[0 0 0], len(arr)=3, cap(arr)=5
	arr1 := arr[1:3]
	arr2 := arr[1:3]
	fmt.Printf("arr=%v, len(arr)=%v, cap(arr)=%v\n", arr, len(arr), cap(arr))
	fmt.Printf("arr1=%v, len(arr1)=%v, cap(arr1)=%v\n", arr1, len(arr1), cap(arr1))

	arr1[0] = 100
	arr1[1] = 200
	arr1 = append(arr1, 300)
	fmt.Printf("arr=%v, len(arr)=%v, cap(arr)=%v\n", arr, len(arr), cap(arr))
	fmt.Printf("arr1=%v, len(arr1)=%v, cap(arr1)=%v\n", arr1, len(arr1), cap(arr1))

	arr2 = append(arr1, 1000, 2000, 3000, 4000)
	fmt.Printf("arr=%v, len(arr)=%v, cap(arr)=%v\n", arr, len(arr), cap(arr))
	fmt.Printf("arr2%v, len(arr2)=%v, cap(arr2)=%v\n", arr2, len(arr2), cap(arr2))
	var s []int
	fmt.Printf("s=%v, len(s)=%v, cap(s)=%v\n", s, len(s), cap(s)) //s=[], len(s)=0, cap(s)=0
	fmt.Printf("s == nil: %v\n", s == nil)                        // true

	// fmt.Printf("s[0]=%v\n", s[0])          // panic: runtime error: index out of range [0] with length 0

	s1 := make([]int, 3)
	fmt.Printf("s1=%v, len(s1)=%v, cap(s1)=%v\n", s1, len(s1), cap(s1)) //s1=[   ], len(s1)=3, cap(s1)=3
	fmt.Printf("s1[0]=%v\n", s1[0])                                     //s1[0]=0

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

	dst2 := make([]int, len(s2)-1)
	n := copy(dst2, s2)
	fmt.Printf("dst2 = %v, 复制数量=%d, 源切片长度=%d\n", dst2, n, len(s2)) // dst2 = [0 1 2 3], 复制数量=4, 源切片长度=5

	dst2[1] = 2
	fmt.Printf("dst2 = %v, s2 = %v, \n", dst2, s2)

}
