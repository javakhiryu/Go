package main

import (
	"fmt"      // Пакет для форматированного ввода-вывода
	"net/http" // Пакет для работы с HTTP-запросами
)

// Функция-обработчик для пути "/hello"
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

// Функция-обработчик для пути "/headers"
func headers(w http.ResponseWriter, r *http.Request) {
	// Перебор всех заголовков запроса
	for name, headers := range r.Header {
		// Перебор всех значений каждого заголовка
		for _, h := range headers {
			// Запись заголовка и его значения в тело ответа
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {
	// Назначение функции hello для пути "/hello"
	http.HandleFunc("/hello", sayHello)

	// Назначение функции headers для пути "/headers"
	http.HandleFunc("/headers", headers)

	// Запуск HTTP-сервера на порту 8080
	http.ListenAndServe(":8080", nil)
}
