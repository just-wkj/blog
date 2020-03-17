package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *sqlx.DB

func initDB() (err error) {
	//数据库信息
	dsn := "root:wangkeji@tcp(127.0.0.1:3306)/sql_test?charset=utf8"
	//连接数据库
	// 不使用 := 赋值,否则都是局部变量
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}

	//可以设置最大连接数
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	initDB()

	sql := `SELECT id, name, age from user where id = 1`
	var u1 user
	db.Get(&u1, sql)
	fmt.Printf("u:%#v\n", u1)

	//var userList = make([]user,0, 10)
	var userList []user
	sql = `SELECT id, name, age from user`
	err := db.Select(&userList, sql)
	if err != nil {
		fmt.Printf("Err:%#v", err)
	}
	fmt.Printf("userList:%#v\n", userList)
}
