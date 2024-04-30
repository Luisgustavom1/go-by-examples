package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		// simulate an expensive operation
		result := j * 2
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- result
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
