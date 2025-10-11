//题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，
//同时统计每个任务的执行时间。

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type TaskScheduler struct {
	Tasks [](func(i int))
	wg    sync.WaitGroup
}

func (s *TaskScheduler) AddTask(v func(sec int)) {
	s.Tasks = append(s.Tasks, v)
	s.wg.Add(1)
}

func (s *TaskScheduler) Run() {
	for len(s.Tasks) > 0 {
		task := s.Tasks[0]
		s.Tasks = s.Tasks[1:len(s.Tasks)]
		go func(v func(sec int)) {
			defer s.wg.Done()
			sec := rand.Intn(1001)
			start := time.Now()
			v(sec)
			cost := time.Since(start)
			fmt.Printf("Cost:%v Expected:%v\n", cost.Milliseconds(), sec)
		}(task)
	}
	s.wg.Wait()
}

func Print(sec int) {
	//sec := rand.Intn(1001)
	time.Sleep(time.Duration(sec) * time.Millisecond)
}
func main() {
	task_scheduler := &TaskScheduler{}
	for i := 0; i < 20; i++ {
		task_scheduler.AddTask(Print)
	}
	task_scheduler.Run()
}
