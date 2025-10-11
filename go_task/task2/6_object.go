//使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
//再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
//为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。

package main

import "fmt"

type Person struct {
	Name string
	Age  int32
}

type Employee struct {
	EmployeeID int32
	P          Person
}

func (E *Employee) PrintInfo() {
	fmt.Printf("name: %v, age: %v, id:%v\n", E.P.Name, E.P.Age, E.EmployeeID)
}

func main() {
	p := Employee{
		EmployeeID: 0,
		P: Person{
			Name: "Tom",
			Age:  10,
		},
	}
	p.PrintInfo()
}
