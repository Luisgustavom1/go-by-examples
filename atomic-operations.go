package main

import (
	"sync"
	"sync/atomic"
)

func main() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}

			defer wg.Done()
		}()
	}

	wg.Wait()

	opsFinal := ops.Load()
	println("ops:", opsFinal)
}
