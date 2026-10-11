package main

import (
	"fmt"
	"time"
)

func main() {
	var Count = 25
	request := make(chan int, Count)

	for i := 1; i <= Count; i++ {
		request <- i
	}
	close(request)

	tick := time.Tick(200 * time.Millisecond)

	for requestPerSecond := range request {
		<-tick

		fmt.Println(" Обработан запрос ", time.Now(), requestPerSecond)
	}
}
