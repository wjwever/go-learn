package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)
	carry := 1
	for i := n - 1; i >= 0; i-- {
		bak := digits[i]
		digits[i] = (bak + carry) % 10
		carry = (bak + carry) / 10
		if carry == 0 {
			return digits
		}
	}

	ret := make([]int, n+1)
	ret[0] = 1
	return ret
}

func main() {
	digits1 := []int{1, 2, 3}
	fmt.Printf("%v", digits1)
	fmt.Printf(":")
	fmt.Printf("%v\n", plusOne(digits1))

	digits2 := []int{4, 3, 2, 1}
	fmt.Printf("%v", digits2)
	fmt.Printf(":")
	fmt.Printf("%v\n", plusOne(digits2))

	digits3 := []int{9}
	fmt.Printf("%v", digits3)
	fmt.Printf(":")
	fmt.Printf("%v\n", plusOne(digits3))
}
