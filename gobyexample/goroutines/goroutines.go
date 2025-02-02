//Ваш код демонстрирует использование горутин для выполнения нескольких функций одновременно в Go. Это позволяет функции f и анонимной функции выполняться параллельно с основным потоком.

package main

import (
	"fmt"
	"time"
)

// Функция f выводит три строки с переданным параметром from и индексом
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Вызов функции f напрямую
	f("direct")

	// Вызов функции f в горутине
	go f("gourutine")

	// Вызов анонимной функции в горутине с передачей аргумента
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Ждем 1 секунду для завершения работы горутин
	time.Sleep(time.Second)
	fmt.Println("done")
}


//Ваш код демонстрирует использование горутин для выполнения нескольких функций одновременно в Go. Это позволяет функции f и анонимной функции выполняться параллельно с основным потоком.
//
//Вот ваш код с комментариями на русском языке для лучшего понимания:


//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//// Функция f выводит три строки с переданным параметром from и индексом
//func f(from string) {
//	for i := 0; i < 3; i++ {
//		fmt.Println(from, ":", i)
//	}
//}
//
//func main() {
//	// Вызов функции f напрямую
//	f("direct")
//
//	// Вызов функции f в горутине
//	go f("gourutine")
//
//	// Вызов анонимной функции в горутине с передачей аргумента
//	go func(msg string) {
//		fmt.Println(msg)
//	}("going")
//
//	// Ждем 1 секунду для завершения работы горутин
//	time.Sleep(time.Second)
//	fmt.Println("done")
//}

//Объяснение:
//
//Функция f: Эта функция выводит три строки, каждая из которых содержит переданную строку from и индекс i.
//Вызов функции f напрямую: Вызов функции f с аргументом "direct" выполняется в основном потоке.
//Вызов функции f в горутине: Вызов функции f с аргументом "gourutine" выполняется в новой горутине, позволяя ей выполняться параллельно с другими задачами.
//Вызов анонимной функции в горутине: Анонимная функция, которая принимает строку msg и выводит её, вызывается в новой горутине с аргументом "going".
//Ожидание завершения работы горутин: Функция time.Sleep(time.Second) приостанавливает выполнение основной программы на 1 секунду, чтобы дать горутинам время завершить свою работу перед завершением программы.
//
//Использование горутин позволяет выполнять задачи параллельно, что может быть полезно для выполнения асинхронных операций, таких как обработка сетевых запросов или выполнение длительных вычислений без блокировки основного потока программы.