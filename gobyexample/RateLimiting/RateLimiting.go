//Ваш код демонстрирует использование таймеров для ограничения скорости обработки запросов в Go. Используется простой limiter для ограничения скорости выполнения запросов и burstyLimiter для разрешения "пакетной" обработки запросов.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем канал для запросов и заполняем его
	requests := make(chan int, 5)
	for i := 1; i < 6; i++ {
		requests <- i
	}
	close(requests)

	// Создаем простой лимитер, который срабатывает каждые 200 миллисекунд
	limiter := time.Tick(200 * time.Millisecond)

	// Обрабатываем запросы с использованием лимитера
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// Создаем "пакетный" лимитер с буфером на 3 события
	burstyLimiter := make(chan time.Time, 3)

	// Заполняем буфер burstyLimiter текущим временем
	for i := 1; i < 4; i++ {
		burstyLimiter <- time.Now()
	}

	// Запускаем горутину для добавления текущего времени в burstyLimiter каждые 200 миллисекунд
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Создаем канал для "пакетных" запросов и заполняем его
	burstyRequests := make(chan int, 5)
	for i := 1; i < 6; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// Обрабатываем "пакетные" запросы с использованием burstyLimiter
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("burstyRequest", req, time.Now())
	}

}

//Объяснение:
//
//Создание канала запросов: Канал requests заполняется пятью запросами и затем закрывается.
//Создание простого лимитера: limiter используется для ограничения скорости обработки запросов, срабатывая каждые 200 миллисекунд. Запросы обрабатываются в цикле for range с использованием лимитера.
//Создание "пакетного" лимитера: burstyLimiter позволяет обрабатывать до трех запросов без задержки. Изначально он заполняется тремя текущими временными метками.
//Горутина для пополнения burstyLimiter: Горутина добавляет текущее время в burstyLimiter каждые 200 миллисекунд, обеспечивая возможность пакетной обработки запросов.
//Создание канала для "пакетных" запросов: Канал burstyRequests заполняется пятью запросами и затем закрывается.
//Обработка "пакетных" запросов: Запросы обрабатываются в цикле for range с использованием burstyLimiter.
//
//Этот пример показывает, как можно реализовать управление скоростью выполнения запросов и пакетную обработку в Go, используя таймеры и каналы.
