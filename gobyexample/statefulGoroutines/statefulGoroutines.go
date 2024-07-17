//В этом примере создаются несколько Go-рутину, которые выполняют операции чтения и записи в карту через каналы. Код также использует атомарные операции для подсчета количества операций чтения и записи.

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

	// Создаем каналы для операций чтения и записи
	reads := make(chan ReadOp)
	writes := make(chan WriteOp)

	// Горутин для обработки операций чтения и записи
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

	// Создаем 100 горутин для выполнения операций чтения
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := ReadOp{
					key:  rand.Intn(5), // Генерируем случайный ключ для чтения
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&ReadOps, 1) // Увеличиваем счетчик операций чтения
				time.Sleep(time.Millisecond)  // Задержка для имитации работы

			}
		}()
	}

	// Создаем 10 горутин для выполнения операций записи
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := WriteOp{
					key:  rand.Intn(5),   // Генерируем случайный ключ для записи
					val:  rand.Intn(100), // Генерируем случайное значение для записи
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&WriteOps, 1) // Увеличиваем счетчик операций записи
				time.Sleep(time.Millisecond)   // Задержка для имитации работы
			}
		}()
	}
	// Ожидание некоторое время для завершения горутин
	time.Sleep(2 * time.Second) // Обновлено для того, чтобы дать время горутинам для выполнения

	// Выводим количество операций
	readRes := atomic.LoadUint64(&ReadOps)
	fmt.Println("Read counter:", readRes)
	writeRes := atomic.LoadUint64(&WriteOps)
	fmt.Println("Write counter:", writeRes)

}
