package main

import "fmt"

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
func double(arr []int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

func main() {
	arr := []int{0, 1, 2, 3}
	fmt.Printf("before: %v\n", arr)
	double(arr)
	fmt.Printf("after: %v\n", arr)
}
