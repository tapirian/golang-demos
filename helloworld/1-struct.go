package main

import (
	"encoding/json"
	"fmt"
	"time"
)

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
	zhangsan = &Person{
		name: "张三",
		sex:  1,
	}
	zhangsan.age = 12
	fmt.Printf("zhangsan : %+v\n", zhangsan)   // &{{ 0} 张三 1}
	fmt.Printf("*zhangsan : %+v\n", *zhangsan) // {{ 0} 张三 1}

	lisi := &Person{
		name: "李四",
		sex:  2,
	}
	fmt.Printf("lisi : %v\n", lisi)
	(*lisi).name = "李思"
	fmt.Printf("lisi : %v\n", lisi)
	lisi.name = "李思1"
	fmt.Printf("lisi : %v\n", lisi)

	wangwu := &Person{
		name: "wangwu",
		sex:  1,
	}
	fmt.Printf("wangwu : %v\n", wangwu)
	// updateName(*wangwu)
	updateName2(wangwu)
	fmt.Printf("name-or : %v\n", wangwu)

	zhangliu := &Person{}
	zhangliu.age = 12
	zhangliu.name = "123"
	zhangliu.sex = 1
	fmt.Printf("zhangliu : %v\n", zhangliu)

	test2()
}

func updateName(p Person) {
	p.name = "蓝星——" + p.name
	fmt.Printf("updateName : %v\n", p)
}

func updateName2(p *Person) {
	p.name = "蓝星——" + p.name
	fmt.Printf("updateName2 : %v\n", p)
}

func test2() {
	type User struct {
		Name       string    `json:"name"`
		Age        uint32    `json:"age"`
		UpdateTime time.Time `json:"updateTime"`
	}

	type Response struct {
		User
		UpdateTime string `json:"updateTime"`
	}

	user := User{
		Name:       "张三",
		Age:        12,
		UpdateTime: time.Now(),
	}

	userResponse := Response{
		User:       user,
		UpdateTime: user.UpdateTime.Format(time.DateTime),
	}

	userData, _ := json.Marshal(user)
	userRespData, _ := json.MarshalIndent(userResponse, "", "  ") // 格式化

	fmt.Printf("user=%s\n", userData) //user={"name":"张三","age":12,"updateTime":"2025-10-09T15:16:30.8851019+08:00"}
	fmt.Printf("userResp=%s\n", userRespData)
	/*
		userResp={
		  "name": "张三",
		  "age": 12,
		  "updateTime": "2025-10-09 15:16:30"
		}
	*/
}
