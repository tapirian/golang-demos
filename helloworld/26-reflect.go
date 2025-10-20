package main

import (
	"fmt"
	"reflect"
)

// Golang 反射
func main() {
	StructExample()
}

func basicExample() {
	var a int = 42
	var b any
	b = a
	println(a, b)

	t1 := reflect.TypeOf(a)
	k1 := t1.Kind()
	v1 := reflect.ValueOf(a)

	t2 := reflect.TypeOf(b)
	k2 := t2.Kind()
	v2 := reflect.ValueOf(b)

	fmt.Println(t1, v1, k1)
	fmt.Println(t2, v2, k2)
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) Greet(pre string) string {
	return pre + u.Name
}

func StructExample() {
	u := User{
		Name: "Alice",
		Age:  30,
	}

	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)

	fmt.Println("Type:", t)
	fmt.Println("Value:", v)

	// 遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		jsonTag := field.Tag.Get("json")
		fmt.Printf("Field Name: %s, Field Type: %s, Field Value: %v, jsonTag: %s\n", field.Name, field.Type, value.Interface(), jsonTag)
	}

	// 调用方法通过反射
	v = reflect.ValueOf(&u) // 注意这里需要传入指针以调用方法
	// 准备参数
	args := []reflect.Value{reflect.ValueOf("Hello, ")}
	// 调用方法
	results := v.MethodByName("Greet").Call(args)
	// 获取返回值
	greeting := results[0].Interface().(string)
	fmt.Println("Greeting:", greeting)

	// 也可以直接调用
	v.MethodByName("Greet").Call([]reflect.Value{reflect.ValueOf("Hello, ")})
}
