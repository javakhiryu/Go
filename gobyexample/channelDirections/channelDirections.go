//Функция ping: Эта функция принимает канал pings для отправки сообщений (chan<- означает "только отправка") и сообщение msg. Она отправляет msg в канал pings.

//Функция pong: Эта функция принимает канал pings для получения сообщений (<-chan означает "только получение") и канал pongs для отправки сообщений (chan<-).
//Она получает сообщение из канала pings и отправляет его в канал pongs.

// Главная функция main:
// Создает два канала pings и pongs с буфером размером 1.
// Вызывает функцию ping, чтобы отправить сообщение "passed message" в канал pings.
// Вызывает функцию pong, чтобы получить сообщение из канала pings и отправить его в канал pongs.
// Читает и выводит сообщение из канала pongs.
package main

import (
	"fmt"
)

// Функция ping отправляет сообщение в канал pings
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Функция pong получает сообщение из канала pings и отправляет его в канал pongs
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	// Создаем два канала с буфером размером 1
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// Отправляем сообщение "passed message" в канал pings
	ping(pings, "passed message")

	// Получаем сообщение из канала pings и отправляем его в канал pongs
	pong(pings, pongs)

	// Читаем и выводим сообщение из канала pongs
	fmt.Println(<-pongs)
}
