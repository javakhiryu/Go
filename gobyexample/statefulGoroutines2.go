package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	Count int
}

func increment(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 1000; i++ {
		c.mu.Lock()
		c.Count++
		c.mu.Unlock()
	}
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go increment(&counter, &wg)
	}

	wg.Wait()

	fmt.Println("Final Count:", counter.Count)
}
