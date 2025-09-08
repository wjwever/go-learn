package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	id := 0

	for i := 1; i < len(nums); i++ {
		if nums[id] != nums[i] {
			id++
			nums[id] = nums[i]
		}
	}

	return id + 1
}

func main() {
	case1 := []int{1, 1, 2}
	fmt.Printf("%v : %v\n", case1, removeDuplicates(case1))

	case2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Printf("%v : %v\n", case2, removeDuplicates(case2))
}
