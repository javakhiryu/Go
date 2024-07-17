package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// apiHandler отвечает на запросы к /api, возвращая "Hello world!".
func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

// healthCheck отвечает на запросы к /healthCheck, возвращая "ok".
func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

// setupHandler регистрирует обработчики для маршрутов /api и /healthCheck в http.ServeMux.
func setupHandler(mux *http.ServeMux) {
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/healthCheck", healthCheck)
}

func main() {
	// Получаем адрес и порт для прослушивания из переменной окружения LISTEN_ADDR.
	// Если переменная не установлена, используем порт по умолчанию :8080.
	listenAddr := os.Getenv("LISTEN_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	// Создаем новый http.ServeMux для маршрутизации запросов.
	mux := http.NewServeMux()

	// Настраиваем маршруты с помощью setupHandler.
	setupHandler(mux)

	// Запускаем сервер и слушаем на заданном адресе.
	// Если сервер не удается запустить, логируем ошибку и завершаем выполнение.
	log.Fatal(http.ListenAndServe(listenAddr, mux))
}
