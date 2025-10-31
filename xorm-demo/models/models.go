package models

type User struct {
	Id      int    `xorm:"not null pk autoincr INTEGER"`
	UsrName string `xorm:"not null unique TEXT"`
}
