package main
import (
    "fmt"
)

func singleNumber(nums []int) int {
    ret := 0
    for _, val :=range(nums) {
        ret ^= val
    }
    return ret
}

func main() {
    // single number 
    nums1 := []int{2, 2, 1}
    fmt.Printf("single number %v : %v\n", nums1, singleNumber(nums1))
}
