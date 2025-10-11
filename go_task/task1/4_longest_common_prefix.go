package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if i == len(strs[j]) || strs[j][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}

func main() {
	strs1 := []string{"flower", "flow", "flight"}
	fmt.Printf("%v : %v\n", strs1, longestCommonPrefix((strs1)))
	strs2 := []string{"dog", "racecar", "car"}
	fmt.Printf("%v : %v\n", strs2, longestCommonPrefix((strs2)))
}
