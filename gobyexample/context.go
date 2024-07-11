package main

import (
	"fmt"      // Пакет для форматированного ввода-вывода
	"net/http" // Пакет для работы с HTTP-запросами
	"time"     // Пакет для работы с временем
)

// Функция-обработчик для пути "/hello"
func hello(w http.ResponseWriter, req *http.Request) {

	// Получение контекста запроса
	ctx := req.Context()
	fmt.Println("server: hello handler started")

	// Сообщение о завершении работы обработчика
	defer fmt.Println("server: hello handler ended")

	select {
	// Ожидание 10 секунд перед отправкой ответа
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")

		// Обработка завершения контекста (например, клиент закрыл соединение)
	case <-ctx.Done():
		// Получение ошибки из контекста
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		//Отправка ошибки клиенту
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	// Назначение функции hello для пути "/hello"
	http.HandleFunc("/hello", hello)
	// Запуск HTTP-сервера на порту 8080
	http.ListenAndServe(":8080", nil)
}
