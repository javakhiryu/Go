//Этот пример демонстрирует, как параллельно отправлять HTTP-запросы и собирать результаты, используя горутины и каналы в Go. Основные компоненты программы включают тип данных result для хранения результатов запросов, функцию get для выполнения HTTP-запросов и основную функцию main, которая управляет запуском горутин и сбором результатов.

package main

import (
	"log"
	"time"
	"net/http"
)

// result представляет результат HTTP-запроса
type result struct {
	url     string 			// URL, к которому был сделан запрос
	err     error 			// Ошибка, если произошла
	latency time.Duration 	// Время задержки запроса

// get выполняет HTTP-запрос к указанному URL и отправляет результат в канал
func get(url string, ch chan <- result){
	start := time.Now() // Запоминаем время начала запроса

	resp, err := http.Get(url) // Выполняем HTTP-запрос
	if err !=nil {
		ch <- result {url, err, 0} // Если произошла ошибка, отправляем результат с ошибкой и нулевой задержкой
	} else{
		t:= time.Since(start).Round(time.Millisecond) // Вычисляем время задержки
		ch <- result {url, nil, t}   // Отправляем результат с задержкой
		resp.Body.Close()   // Закрываем тело ответа
	}
}

func main() {
	results:=make(chan result) // Создаем канал для передачи результатов
	list := []string{ // Список URL для запросов

		"https://google.com",
		"https://nytimes.com",
		"https://amazon.com",
		"https://wsj.com",

	}
	for _, url := range list{ // Запускаем горутину для каждого URL
		go get(url, results)
	}

	for range list{ // Собираем результаты
		r := <-results // Получаем результат из канала

		if r.err !=nil{ // Если произошла ошибка
			log.Printf("%-20s %s\n", r.url, r.err)
		} else { // Если запрос успешен
			log.Printf("%-20s %s\n", r.url, r.latency)
		}
	}
}

//Этот пример показывает, как эффективно использовать горутины и каналы для параллельного выполнения задач и сбора результатов.