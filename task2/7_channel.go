package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func produce(sig chan int, wg *sync.WaitGroup) {
	for i := 1; i <= 10; i++ {
		sig <- i
	}
	close(sig)
	wg.Done()
}

func consume(sig chan int, wg *sync.WaitGroup) {
	for v := range sig {
		fmt.Println(v)
	}
	wg.Done()
}

func main() {
	sig := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go produce(sig, &wg)
	go consume(sig, &wg)
	wg.Wait()
}
