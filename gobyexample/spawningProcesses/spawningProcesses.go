package main

import (
	"fmt"     // Пакет для форматированного вывода
	"io"      // Пакет для работы с вводом-выводом
	"os/exec" // Пакет для выполнения внешних команд
)

func main() {
	// Создаем команду для выполнения команды "date"
	dateCmd := exec.Command("date")

	// Выполняем команду и получаем её вывод
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err) // Если произошла ошибка, прерываем выполнение программы
	}

	fmt.Println("> date")
	fmt.Println(string(dateOut)) // Выводим результат команды "date"

	// Пробуем выполнить команду "date" с несуществующим флагом "-x"
	_, err = exec.Command("date", "-x").Output()
	if err != nil {
		// Обрабатываем ошибку в зависимости от её типа
		switch e := err.(type) {
		case *exec.Error:
			fmt.Println("failed executing:", err)

		case *exec.ExitError:
			fmt.Println("command exit rc=", e.ExitCode())

		default:
			panic(err)
		}
	}

	// Создаем команду для выполнения команды "grep hello"
	grepCmd := exec.Command("grep", "hello")

	// Получаем канал для записи ввода в команду "grep"
	grepIn, _ := grepCmd.StdinPipe()

	// Получаем канал для записи ввода в команду "grep"
	grepOut, _ := grepCmd.StdoutPipe()

	// Запускаем команду
	grepCmd.Start()

	// Пишем данные на вход команды "grep"
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	// Читаем результат выполнения команды "grep"
	grepBytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait() // Ждем завершения команды

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes)) // Выводим результат команды "grep"

	// Создаем команду для выполнения сложной команды "ls -a -l -h" через оболочку
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	// Выполняем команду и получаем её вывод
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut)) // Выводим результат команды "ls"
}

//Выполнение команд оболочки из программы на Go может быть полезным в различных сценариях. Вот несколько примеров и причин, по которым это может понадобиться:

//1. Автоматизация задач
//Выполнение команд оболочки позволяет автоматизировать задачи, которые обычно выполняются вручную в командной строке. Например:
//
//Запуск резервного копирования с помощью rsync.
//Сборка проекта с помощью make.
//Выполнение скриптов для развертывания и конфигурации серверов.

//package main
//
//import (
//    "fmt"
//    "os/exec"
//)
//
//func main() {
//    // Запуск команды ls для получения списка файлов в текущей директории
//    cmd := exec.Command("ls")
//    output, err := cmd.Output()
//    if err != nil {
//        panic(err)
//    }
//    fmt.Println(string(output))
//}

//2. Интеграция с существующими инструментами
//Иногда полезно использовать уже существующие команды и инструменты, вместо того чтобы писать новый код для выполнения тех же задач. Это позволяет интегрировать вашу программу с другими утилитами и инструментами. Например:
//
//Использование grep для фильтрации текстовых данных.
//Запуск команды tar для архивации файлов.
//Использование curl для выполнения HTTP-запросов.

//package main
//
//import (
//    "fmt"
//    "io"
//    "os/exec"
//)
//
//func main() {
//    // Использование команды grep для поиска строки в тексте
//    cmd := exec.Command("grep", "hello")
//    stdin, _ := cmd.StdinPipe()
//    stdout, _ := cmd.StdoutPipe()
//    cmd.Start()
//    io.WriteString(stdin, "hello world\nbye world")
//    stdin.Close()
//    output, _ := io.ReadAll(stdout)
//    cmd.Wait()
//    fmt.Println(string(output))
//}

//3. Мониторинг и сбор метрик
//Вы можете запускать команды для получения системной информации и метрик. Например:
//
//Запуск команды df для получения информации о свободном дисковом пространстве.
//Использование top или ps для мониторинга использования ресурсов.

//package main
//
//import (
//    "fmt"
//    "os/exec"
//)
//
//func main() {
//    // Получение информации о свободном дисковом пространстве
//    cmd := exec.Command("df", "-h")
//    output, err := cmd.Output()
//    if err != nil {
//        panic(err)
//    }
//    fmt.Println(string(output))
//}

//4. Обработка пользовательских запросов
//Программы могут выполнять команды оболочки в ответ на действия пользователя. Например:
//
//Запуск определенной команды при получении HTTP-запроса на веб-сервере.
//Выполнение команды на основе ввода пользователя в интерфейсе командной строки.

//package main
//
//import (
//    "fmt"
//    "net/http"
//    "os/exec"
//)
//
//func handler(w http.ResponseWriter, r *http.Request) {
//    cmd := exec.Command("date")
//    output, err := cmd.Output()
//    if err != nil {
//        fmt.Fprintf(w, "Error: %v", err)
//        return
//    }
//    fmt.Fprintf(w, "Current date and time: %s", output)
//}
//
//func main() {
//    http.HandleFunc("/date", handler)
//    http.ListenAndServe(":8080", nil)
//}
