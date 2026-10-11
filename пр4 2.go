package main

import (
	"fmt"
	"sync"
)

func worker(id int, tasks <-chan int, out chan<- int, group *sync.WaitGroup) {
	defer group.Done()
	for t := range tasks {
		out <- t * t
	}
}

func main() {
	tasks := make(chan int, 10)
	out := make(chan int, 10)
	var group sync.WaitGroup

	for i := 1; i <= 3; i++ {
		group.Add(1)
		go worker(i, tasks, out, &group)
	}

	for n := 1; n <= 10; n++ {
		tasks <- n
	}
	close(tasks)
	go func() {
		group.Wait()
		close(out)
	}()
	for val := range out {
		fmt.Println("Результат:", val)
	}
}
