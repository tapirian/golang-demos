package main

import (
	"flag"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
}

var currentUser User = User{
	Name: "张三",
	Age:  18,
}

func main() {
	var name string
	var age int
	// 创建新的标志集
	setCmd := flag.NewFlagSet("setUser", flag.ExitOnError)
	setCmd.StringVar(&name, "name", "", "姓名")
	setCmd.IntVar(&age, "age", 0, "年龄")

	getCmd := flag.NewFlagSet("getUser", flag.ExitOnError)

	// 解析命令行参数
	if os.Args[1] == "setUser" {
		setCmd.Parse(os.Args[2:])
		currentUser.Name = name
		currentUser.Age = age
		fmt.Printf("setUser Success: 姓名:%s \t 年龄:%d\n", currentUser.Name, currentUser.Age)
	} else if os.Args[1] == "getUser" {
		getCmd.Parse(os.Args[2:])
		fmt.Printf("getUser: 姓名:%s \t 年龄:%d\n", currentUser.Name, currentUser.Age)
	} else {
		fmt.Println("无效的命令")
		return
	}
}
