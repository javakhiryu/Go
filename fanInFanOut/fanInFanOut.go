//Этот код на Go демонстрирует использование каналов, горутин и концепций "fan-out" и "fan-in" для эффективного параллельного поиска простых чисел. Давайте рассмотрим каждую часть кода и добавим комментарии для лучшего понимания.

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// faIn объединяет несколько входных каналов в один выходной канал.
func faIn[T any](done <-chan int, channels ...<-chan T) <-chan T {
	var wg sync.WaitGroup
	fannedInStream := make(chan T)

	transfer := func(c <-chan T) {
		defer wg.Done()
		for i := range c {
			select {
			case <-done:
				return
			case fannedInStream <- i:
			}
		}
	}

	// Запуск горутины для каждого канала.
	for _, c := range channels {
		wg.Add(1)
		go transfer(c)
	}

	// Закрытие выходного канала после завершения всех горутин.
	go func() {
		wg.Wait()
		close(fannedInStream)
	}()
	return fannedInStream
}

// repeatFunc запускает функцию fn в отдельной горутине и отправляет её результат в канал stream до тех пор,
// пока не получит сигнал завершения из канала done.
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

// take читает из канала stream до n раз или до получения сигнала завершения из канала done.
func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)

	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:

			}
		}
	}()
	return taken
}

// primeFinder получает случайные числа из randomIntStream и проверяет, являются ли они простыми.
func primeFinder(done <-chan int, randomIntStream <-chan int) <-chan int {
	isPrime := func(randomInt int) bool {
		for i := randomInt - 1; i > 1; i-- {
			if randomInt%i == 0 {
				return false
			}
		}
		return true
	}
	primes := make(chan int)
	go func() {
		defer close(primes)
		for {
			select {
			case <-done:
				return
			case randomInt := <-randomIntStream:
				if isPrime(randomInt) {
					primes <- randomInt
				}
			}
		}
	}()
	return primes
}

func main() {
	start := time.Now()
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int { return rand.Intn(50000000) }
	randIntStream := repeatFunc(done, randNumFetcher)

	// Fan-out: создаем несколько горутин для поиска простых чисел, равное количеству ядер CPU.
	CPUCount := runtime.NumCPU()
	primeFinderChannels := make([]<-chan int, CPUCount)
	for i := 0; i < CPUCount; i++ {
		primeFinderChannels[i] = primeFinder(done, randIntStream)
	}

	// Fan-in: объединяем результаты из всех горутин в один канал.
	fannedInStream := faIn(done, primeFinderChannels...)

	// Выводим первые 10 простых чисел.
	for random := range take(done, fannedInStream, 10) {
		fmt.Println(random)
	}
	fmt.Println(time.Since(start))
}

//Функция faIn:
//
//Объединяет несколько входных каналов в один выходной канал.
//Использует sync.WaitGroup для отслеживания завершения всех горутин.
//Закрывает выходной канал после завершения всех горутин.
//Функция repeatFunc:
//
//Генерирует значения с помощью функции fn и отправляет их в канал stream до тех пор, пока не получит сигнал завершения из канала done.
//Функция take:
//
//Читает значения из канала stream до тех пор, пока не будет прочитано n значений или не будет получен сигнал завершения из канала done.
//Функция primeFinder:
//
//Получает случайные числа из randomIntStream и проверяет, являются ли они простыми.
//Отправляет простые числа в канал primes.
//Функция main:
//
//Создает канал done для отправки сигнала завершения.
//Создает канал randIntStream для генерации случайных чисел.
//Использует "fan-out" для запуска нескольких горутин, которые ищут простые числа.
//Использует "fan-in" для объединения результатов из всех горутин в один канал.
//Выводит первые 10 простых чисел и измеряет время выполнения.
//Этот пример демонстрирует, как можно эффективно использовать параллельные вычисления в Go для решения задач, требующих большого объема вычислений.
