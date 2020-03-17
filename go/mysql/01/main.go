package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

//Go 连接mysql

var db *sql.DB

func initDB() (err error) {
	//数据库信息
	dsn := "root:wangkeji@tcp(127.0.0.1:3306)/sql_test?charset=utf8"
	//连接数据库
	// 不使用 := 赋值,否则都是局部变量
	db, err = sql.Open("mysql", dsn)
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
func query() {

}

func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("db err:$v\n", err)
		return
	}
	fmt.Println("连接数据库成功")

	queryOneDemo(1)
	queryOneDemo(2)

	queryInsert("name1", 22)
	queryAll()

}

type user struct {
	id   int
	name string
	age  int
}

func queryOneDemo(id int) {
	var user1 user
	sql := "SELECT id,name,age FROM user where id=?"
	rowObj := db.QueryRow(sql, id)
	//必须调用Scan, 调用Scan才可以自动释放连接
	rowObj.Scan(&user1.id, &user1.name, &user1.age)

	fmt.Printf("user1: %#v\n", user1)
}

func queryAll() {
	fmt.Println("queryall ===\n")
	sql := "SELECT id,name,age FROM user where id > 0"

	rows, err := db.Query(sql)
	if err != nil {
		fmt.Printf("查询失败:%#v\n", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user1 user
		err = rows.Scan(&user1.id, &user1.name, &user1.age)
		if err != nil {
			fmt.Printf("rows查询失败:%#v\n", err)
			return
		}
		fmt.Printf("user1: %#v\n", user1)
	}
	err = rows.Err() // get any error encountered during iteration

}
func queryInsert(name string, age int) {
	sql := `INSERT INTO user(name, age) values(?,?),("xx",2)`
	ret, err := db.Exec(sql, name, age)
	if err != nil {
		fmt.Printf("插入:%#v\n", err)
		return
	}
	//插入多行ID结果不正确,仅可计算单行
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("插入2:%#v\n", err)
		return
	}
	fmt.Printf("插入ID:%d\n", id)

	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("插入2:%#v\n", err)
		return
	}
	fmt.Printf("影响行数:%d\n", num)

}
