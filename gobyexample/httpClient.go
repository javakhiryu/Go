package main

import (
	"bufio"    // Пакет для работы с буферами ввода-вывода
	"fmt"      // Пакет для форматированного ввода-вывода
	"net/http" // Пакет для работы с HTTP-запросами
)

func main() {

	// Отправка GET-запроса на указанный URL
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		// Обработка ошибки при выполнении запроса
		panic(err)
	}

	// Закрытие тела ответа после завершения работы с ним
	defer resp.Body.Close()

	// Вывод статуса ответа
	fmt.Println("Response status:", resp.Status)

	// Создание сканера для чтения тела ответа построчно
	scanner := bufio.NewScanner(resp.Body)

	// Чтение первых 5 строк из тела ответа
	for i := 0; scanner.Scan() && i < 5; i++ {
		// Вывод каждой строки
		fmt.Println(scanner.Text())
	}
	// Проверка ошибок при сканировании
	if err := scanner.Err(); err != nil {
		// Обработка ошибки при сканировании
		panic(err)
	}
}

//Этот код отправляет запрос на сайт https://gobyexample.com, получает ответ и выводит первые 5 строк из тела ответа. Он также выводит статус ответа и обрабатывает возможные ошибки.