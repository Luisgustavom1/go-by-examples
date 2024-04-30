package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	c.mu.Lock()
	c.counters[name]++
	c.mu.Unlock()
}

func main() {
	c := Container{
		mu: sync.Mutex{},
		counters: map[string]int{
			"a": 0, "b": 0,
		},
	}

	var wg sync.WaitGroup

	incCounter := func(name string, n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			c.inc(name)
		}
	}

	wg.Add(3)
	go incCounter("a", 1000)
	go incCounter("b", 1000)
	go incCounter("a", 1000)

	wg.Wait()
	fmt.Println(c.counters)
}
