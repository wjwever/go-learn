package main

import (
	"fmt"
	"log"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
type Employee struct {
	id         int    `db:"id"`
	name       string `db:"name"`
	department string `db:"department"`
	salary     int    `db:"salary"`
}
*/

type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`
	Department string `db:"department"`
	Salary     int    `db:"salary"`
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

	createSql := `create table if not exists employees (
		id int primary key auto_increment,
		name varchar(30) not null,
		department varchar(30) not null,
		salary int not null
	)`

	_, err = db.Exec(createSql)
	checkErr(err)

	_, err = db.Exec("truncate table employees")
	checkErr(err)

	insertSql := `insert into employees values(?, ?, ?, ?)`
	_, err = db.Exec(insertSql, 0, "张三", "技术部", 20000)
	checkErr(err)

	_, err = db.Exec(insertSql, 0, "王二", "技术部", 30000)
	checkErr(err)

	// 查询技术部人员
	var persons []Employee
	err = db.Select(&persons, "select * from employees where department=?", "技术部")
	checkErr(err)
	for _, person := range persons {
		fmt.Printf("技术部人员 %v\n", person)
	}
	// 查询工资最高的人员
	var maxSalary Employee
	err = db.Get(&maxSalary, "select * from employees order by salary desc limit 1")
	checkErr(err)
	fmt.Printf("工资最高人员： %v\n", maxSalary)
}
