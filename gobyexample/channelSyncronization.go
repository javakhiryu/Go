//Ваш код демонстрирует использование каналов для синхронизации горутин в Go. 
//Ваша программа создает канал done и запускает горутину worker, которая выполняет некоторую работу, а затем сигнализирует о своем завершении, отправляя значение true в канал done. 
//Главная функция main блокируется, ожидая завершения работы горутины, считывая значение из канала done.

package main

import (
	"fmt"
	"time"
)

// Функция worker, которая выполняет некоторую работу и сигнализирует о завершении
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second) // Симуляция работы в течение 1 секунды
	fmt.Print("done")

	done <- true // Отправляем значение в канал, сигнализируя о завершении работы

func main() {
	done := make(chan bool, 1) // Создаем канал с буфером размером 1
	go worker(done) // Запускаем горутину worker
	<-done // Ожидаем завершения работы горутины, считывая значение из каналаv
}

//Функция worker: Эта функция принимает канал done типа chan bool. 
//Она выполняет некоторую работу (симулируемую задержкой на 1 секунду) и затем отправляет значение true в канал done, сигнализируя о завершении работы.

//Главная функция main:
	//Создает канал done с буфером размером 1.
	//Запускает горутину worker, передавая ей канал done.
	//Блокируется, ожидая значение из канала done, что указывает на завершение работы горутины.