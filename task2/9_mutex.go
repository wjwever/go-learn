package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
type Counter struct {
	sync.Mutex
	id int
}

func (c *Counter) Inc(wg *sync.WaitGroup) {
	c.Lock()
	defer c.Unlock()
	defer wg.Done()
	c.id++
}

func (c *Counter) Value() int {
	c.Lock()
	defer c.Unlock()
	return c.id
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{id: 0}
	wg.Add(10 * 1000)
	for i := 0; i < 10; i++ {
		go func() {
			for tmp := 0; tmp < 1000; tmp++ {
				counter.Inc(&wg)
			}
		}()
	}
	wg.Wait()

	fmt.Println(counter.Value())
}
