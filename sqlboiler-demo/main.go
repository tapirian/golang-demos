package main

import (
	"context"
	"database/sql"
	"log"
	models "sqlboiler-demo/my_models"

	"github.com/aarondl/null/v8"
	"github.com/aarondl/sqlboiler/v4/boil"
	"github.com/aarondl/sqlboiler/v4/queries/qm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name  string
	Email string
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("捕获panic: %v", r)
		}
	}()

	dsn := "root:123456@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=true&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 测试数据库连接
	if err := db.Ping(); err != nil {
		panic(err)
	}

	// 设置全局默认数据库连接， 并开启调试
	boil.SetDB(db)
	boil.DebugMode = true

	ctx := context.Background()

	// 创建用户
	user := &models.User{
		Name:  null.NewString("zhangsan", true),
		Email: null.NewString("zhangsan@163.com", true),
	}
	if err := user.Insert(ctx, db, boil.Infer()); err != nil {
		panic(err)
	}
	log.Printf("插入数据： %v \n", *user)

	// 查询用户--根据id
	selectUser, err := models.FindUser(ctx, db, user.ID)
	log.Printf("根据id查用户: %v \n", *selectUser)

	// 查询用户--根据条件
	selectUserByCond, err := models.Users(models.UserWhere.Name.EQ(null.StringFrom("zhangsan")),
		qm.OrderBy(models.UserColumns.ID+" DESC"),
	).All(ctx, db)
	if err != nil {
		panic(err)
	}
	log.Printf("根据条件查询用户： %v \n", selectUserByCond)

	// 更新用户
	user.Name = null.StringFrom("lisi")
	_, err = user.Update(ctx, db, boil.Infer())
	if err != nil {
		panic(err)
	}
	log.Printf("修改name=%s \n", user.Name)

	// 删除用户
	_, err = user.Delete(ctx, db)
	if err != nil {
		panic(err)
	}
	log.Printf("删除用户: %v \n", *user)
}
