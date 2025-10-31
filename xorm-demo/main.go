package main

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
	"xorm.io/xorm"
)

type User struct {
	Id   int64
	Name string `xorm:"varchar(25) notnull unique 'usr_name' comment('姓名')"`
}

func main() {
	// 连接数据库（创建引擎）
	engine, err := xorm.NewEngine("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	user := new(User)
	user.Name = "Putin"

	// 同步结构体到数据库
	err = engine.Sync(new(User))
	if err != nil {
		panic(err)
	}

	// 新增数据
	affected, err := engine.Insert(user)
	if err != nil {
		panic(err)
	}
	println("affected rows: ", affected)
	log.Printf("insert user: %+v \n", *user)

	// 修改数据
	user.Name = "Red"
	_, err = engine.ID(user.Id).Update(user)
	if err != nil {
		panic(err)
	}
	log.Printf("update user: %+v \n", *user)

	// 查询数据
	var selectUser User
	has, err := engine.ID(user.Id).Get(&selectUser)
	if err != nil {
		panic(err)
	}
	if has {
		log.Printf("select user: %+v \n ", selectUser)
	} else {
		println("没有数据")
	}

	// 删除数据
	_, err = engine.ID(user.Id).Delete(new(User))
	if err != nil {
		panic(err)
	}
	println("deleted!!")
}
