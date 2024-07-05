package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type ReadOp struct {
	key  int
	resp chan int
}

type WriteOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	var ReadOps uint64
	var WriteOps uint64
	reads := make(chan ReadOp)
	writes := make(chan WriteOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]

			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := ReadOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&ReadOps, 1)
				time.Sleep(time.Millisecond)

			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := WriteOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&WriteOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	readRes := atomic.LoadUint64(&ReadOps)
	fmt.Println("Read counter:", readRes)
	writeRes := atomic.LoadUint64(&WriteOps)
	fmt.Println("Write counter:", writeRes)

}
