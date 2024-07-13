package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main(){

	// Создаем канал для получения уведомлений о сигналах.
	sigs :=make(chan os.Signal, 1)

	// Уведомляем о сигналах SIGINT (Ctrl+C) и SIGTERM (терминация).
	//Функция signal.Notify регистрирует канал sigs для получения уведомлений о сигналах SIGINT и SIGTERM. SIGINT генерируется при нажатии Ctrl+C, а SIGTERM обычно используется для завершения работы программы.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	 // Создаем канал для уведомления о завершении работы программы.
	done := make(chan bool, 1)

	// Запускаем горутину для обработки сигналов.
	go func ()  {
		// Ждем получения сигнала
		sig :=<-sigs
		fmt.Println()
		fmt.Println(sig)
		 // Отправляем уведомление о завершении работы.
		done <-true
	}()

	fmt.Println("awaiting signal")

	// Блокируем выполнение основного потока до получения уведомления о завершении.
	<-done
	fmt.Println("exiting")
}