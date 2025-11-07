package main

import (
	"context"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 定义数据表结构
type User struct {
	gorm.Model
	Name  string
	Email string
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	ctx := context.Background()

	// 自动迁移模式
	db.AutoMigrate(&User{})

	// 创建
	err = gorm.G[User](db).Create(ctx, &User{Name: "John", Email: "john@gmail.com"})
	if err != nil {
		panic("failed to create user")
	}

	// 读取
	user, err := gorm.G[User](db).Where("name = ?", "John").First(ctx)
	if err != nil {
		panic("failed to read user")
	}
	println("User:", user.Name, user.Email)

	// 更新
	affectRows, err := gorm.G[User](db).Where("name = ?", "John").Updates(ctx, User{Email: "123@gmail.com"})
	if err != nil {
		panic("failed to update user")
	}
	println("Updated rows:", affectRows)

	// 查询
	var users []User
	users, err = gorm.G[User](db).Where("name = ?", "John").Find(ctx)
	if err != nil {
		panic("failed to query users")
	}
	println("Queried users count:", len(users))

	for _, u := range users {
		println("User:", u.Name, u.Email)
	}

	// 删除
	affectRows, err = gorm.G[User](db).Where("name = ?", "John").Delete(ctx)
	if err != nil {
		panic("failed to delete user")
	}
	println("Deleted rows:", affectRows)
}
