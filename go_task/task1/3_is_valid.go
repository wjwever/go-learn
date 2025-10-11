package main

import (
	"fmt"
)

func isValid(s string) bool {
	stk := make([]byte, 0)
	book := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, c := range []byte(s) {
		if book[c] > 0 {
			if len(stk) == 0 || stk[len(stk)-1] != book[c] {
				return false
			}
			stk = stk[:len(stk)-1]
		} else {
			stk = append(stk, c)
		}
	}
	return len(stk) == 0
}

func main() {
	cases := []string{"()", "()[]{}", "(]", "([])", "([)]"}
	for _, v := range cases {
		fmt.Printf("%v    %v\n", v, isValid(v))
	}
}
