package main

import (
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Book struct {
	ID     int     `db:"id"`
	Titile string  `db:"title"`
	Author *string `db:"author"`
	Price  float32 `db:"price"`
}

func checkErr(err error) {
	if err != nil {
		_, _, line, _ := runtime.Caller(1)
		log.Fatalf("error happens: %v \nline:%v", err, line)
	}
}

func main() {
	db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True")
	defer db.Close()
	checkErr(err)

	_, err = db.Exec("create database if not exists task3")
	checkErr(err)

	_, err = db.Exec("use task3")
	checkErr(err)

	_, err = db.Exec("drop table if exists books")
	checkErr(err)

	createSql := `create table if not exists books(
		id int primary key auto_increment,
		title varchar(30) not null,
		author varchar(30) ,
		price decimal(10, 2) not null
	)`

	_, err = db.Exec(createSql)
	checkErr(err)

	_, err = db.Exec("truncate table employees")
	checkErr(err)

	insertSql := `insert into books values(?, ?, ?, ?)`
	_, err = db.Exec(insertSql, 0, "平凡的世界", nil, 50.2)
	checkErr(err)

	_, err = db.Exec(insertSql, 0, "战争与和平", "托尔斯泰", 45.9)
	checkErr(err)

	var books []Book
	err = db.Select(&books, "select * from books where price>?", 50)
	checkErr(err)
	for _, person := range books {
		fmt.Printf("价格大于50的书籍: %v\n", person)
	}
}
