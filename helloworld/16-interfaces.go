package main

import "fmt"

type animaler interface {
	eat() string
	bark() string
}

type dog struct {
	name string
}

type cat struct {
	name string
}

func (c cat) bark() string {
	return c.name + " miaomiao"
}

func (d dog) bark() string {
	return d.name + " wangwang"
}

func (c cat) eat() string {
	return c.name + " 🍔"
}

func (d dog) eat() string {
	return d.name + " 🍪"
}

func main() {
	// var dogbark, catbark string
	var d dog
	d.name = "xiaohei"
	var c cat
	c.name = "xiaobai"

	animalerAction(d)
	animalerAction(c)
}

func animalerAction(a animaler) {
	bark := a.bark()
	fmt.Printf("bark: %v\n", bark)

	eat := a.eat()
	fmt.Printf("eat: %v\n", eat)
}
