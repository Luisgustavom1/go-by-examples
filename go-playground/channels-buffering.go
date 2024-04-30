package main

import (
	"fmt"
	"time"
)

func work(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("Sending ", i)
		ch <- i
		time.Sleep(time.Millisecond * 100)
	}
	close(ch)
}

func main() {
	ch := make(chan int, 3)
	go work(ch)

	for i := range ch {
		fmt.Println("Reading  ", i)
	}
}
