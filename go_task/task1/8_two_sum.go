package main

import "fmt"

func twoSum(nums []int, target int) []int {
	book := make(map[int]int, 0)
	for i, v := range nums {
		if p, ok := book[target-v]; ok {
			return []int{p, i}
		}
		book[v] = i
	}
	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("%v  %v : %v\n", nums, target, twoSum(nums, target))
}
