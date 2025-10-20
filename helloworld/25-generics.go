package main

func main() {
	StackExample()
}

// 基本泛型函数案例 -- 交换两个值
func Swap[T any](a, b T) (T, T) {
	return b, a
}

func basicExamle() {
	a, b := 1, 2
	a, b = Swap(a, b)

	s1, s2 := "hello", "world"
	s1, s2 = Swap(s1, s2)

	println(s1, s2) // Output: world hello
	println(a, b)   // Output: 2 1
}

// 使用约束 -- 排序案例
type Ordered interface {
	~int | ~float64 | ~string
}

func Min[T Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func ConstraintExample() {
	println(Min(3, 5))            // Output: 3
	println(Min(3.5, 2.1))        // Output: 2.1
	println(Min("apple", "pear")) // Output: apple
}

// 结构体泛型案例 -- 栈实现
type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	if len(s.elements) == 0 {
		var zero T
		return zero
	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element
}

func StackExample() {
	var intStack Stack[int]
	intStack.Push(1)
	intStack.Push(2)
	println(intStack.Pop()) // Output: 2
	println(intStack.Pop()) // Output: 1

	var stringStack Stack[string]
	stringStack.Push("hello")
	stringStack.Push("world")
	println(stringStack.Pop()) // Output: world
	println(stringStack.Pop()) // Output: hello
}
