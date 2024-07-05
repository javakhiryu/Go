package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Gourutine with %d started working\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Gourutine with %d ended working\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}
	wg.Wait()
}
