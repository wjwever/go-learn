package main
import (
    "fmt"
)

func isPalindrome(x int) bool {
    if x < 0 || (x % 10 == 0 && x != 0) {
        return false
    }

    tmp := 0
    for x > tmp {
        tmp = tmp * 10 + x % 10
        x /= 10
    }
    return  x == tmp || x == (tmp / 10)
}

func main() {
    fmt.Printf("%v : %v\n", 121, isPalindrome(121))
    fmt.Printf("%v : %v\n", -121, isPalindrome(-121))
    fmt.Printf("%v : %v\n", 10, isPalindrome(10))
}
