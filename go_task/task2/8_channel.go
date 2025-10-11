package main

import "fmt"

//题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。

func produce(sig chan int, tmp chan int) {
	for i := 1; i <= 100; i++ {
		sig <- i
	}
	close(sig)
	tmp <- 1
}

func consume(sig chan int, tmp chan int) {
	for v := range sig {
		fmt.Println(v)
	}
	tmp <- 2
}

func main() {
	sig := make(chan int, 100)
	tmp := make(chan int)
	go produce(sig, tmp)
	go consume(sig, tmp)
	<-tmp
	<-tmp
}
