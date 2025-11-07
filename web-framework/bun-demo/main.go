package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID    int64  `bun:",pk,autoincrement"`
	Name  string `bun:",notnull"`
	Email string `bun:",unique"`
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("捕获panic: ", r)
		}
	}()

	ctx := context.Background()

	// 数据库连接
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:./test.db?cache=shared")
	if err != nil {
		panic(err)
	}
	defer sqldb.Close()

	// 创建Bun实例
	db := bun.NewDB(sqldb, sqlitedialect.New())

	// 查询钩子，打开调试
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	//新建表
	_, err = db.NewCreateTable().Model((*User)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		panic(err)
	}

	// 新增数据
	user := &User{Name: "John Doe", Email: "john@example.com"}
	_, err = db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Insert User: %+v \n ", user)

	// 查询数据
	var selectedUser User
	err = db.NewSelect().Model(&selectedUser).Where("email = ?", "john@example.com").Scan(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Select User: %+v \n ", selectedUser)

	// 修改数据
	updatedUser := &User{Name: "John", Email: "john@163.com"}
	_, err = db.NewUpdate().Model(updatedUser).Where("name = ?", "John Doe").Exec(ctx)
	if err != nil {
		panic(err)
	}

	// 修改之后再查询
	var selectAgainUser User
	err = db.NewSelect().Model(&selectAgainUser).Where("name = ?", "John").Scan(ctx)
	fmt.Printf("Select user again: %+v \n", selectAgainUser)

	// 删除
	_, err = db.NewDelete().Model((*User)(nil)).Where("name = ?", "John").Exec(ctx)
	if err != nil {
		panic(err)
	}

	var deleteQueryUser User
	err = db.NewSelect().Model(&deleteQueryUser).Where("name = ?", "John").Scan(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Printf("after deleted user to query again: %+v \n", deleteQueryUser)
}
