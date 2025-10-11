package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	ret := make([][]int, 0)
	val := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] <= val[1] {
			if val[1] < cur[1] {
				val[1] = cur[1]
			}
		} else {
			ret = append(ret, val)
			val = cur
		}
	}
	ret = append(ret, val)
	return ret
}

func main() {
	intervals1 := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Printf("%v : %v\n", intervals1, merge(intervals1))

	intervals2 := [][]int{{1, 4}, {4, 5}}
	fmt.Printf("%v : %v\n", intervals2, merge(intervals2))
}
