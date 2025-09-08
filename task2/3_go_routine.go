package main

import "fmt"

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func print_odd(done chan bool) {
	for i := 1; i < 10; i += 2 {
		fmt.Println(i)
	}
	done <- true
}

func print_even(done chan bool) {
	for i := 0; i < 10; i += 2 {
		fmt.Println(i)
	}
	done <- true
}

func main() {
	done := make(chan bool, 2)
	go print_odd(done)
	go print_even(done)
	<-done
	<-done
}
