package main

import "fmt"

type Animal struct {
	name string
	age  uint32
}

type Person struct {
	Animal
	name string
	sex  uint32
}

func main() {
	zhangsan := new(Person)
	*zhangsan = Person{
		name: "张三",
		sex:  1,
	}
	fmt.Printf("zhangsan : %v\n", zhangsan)   // &{{ 0} 张三 1}
	fmt.Printf("*zhangsan : %v\n", *zhangsan) // {{ 0} 张三 1}

	lisi := &Person{
		name: "李四",
		sex:  2,
	}
	fmt.Printf("lisi : %v\n", lisi)
	lisi.name = "李思"
	fmt.Printf("lisi : %v\n", lisi)

	wangwu := &Person{
		name: "wangwu",
		sex:  1,
	}
	fmt.Printf("wangwu : %v\n", wangwu)
	updateName2(wangwu)
	fmt.Printf("name-or : %v\n", wangwu)

	zhangliu := &Person{}
	zhangliu.age = 12
	zhangliu.name = "123"
	zhangliu.sex = 1
	fmt.Printf("zhangliu : %v\n", zhangliu)
}

func updateName(p Person) {
	p.name = "蓝星——" + p.name
	fmt.Printf("updateName : %v\n", p)
}

func updateName2(p *Person) {
	p.name = "蓝星——" + p.name
	fmt.Printf("updateName2 : %v\n", p)
}
