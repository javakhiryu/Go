package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {
	// Использование стандартного логгера для записи сообщения
	log.Println("standart logger")

	// Устанавливаем флаги для стандартного логгера: включаем метку времени и микросекунды
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// Устанавливаем флаги для стандартного логгера: включаем метку времени и короткое имя файла/номер строки
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("with file/line")

	// Создаем новый логгер, который будет выводить сообщения с префиксом "my:"
	mylog := log.New(os.Stdout, "my: ", log.LstdFlags)
	mylog.Println("from mylog")

	// Меняем префикс логгера на "ohmy:"
	mylog.SetPrefix("ohmy")
	mylog.Println("from mylog")

	// Создаем буфер для хранения лог-сообщений
	var buf bytes.Buffer

	// Создаем логгер, который будет записывать сообщения в буфер
	buflog := log.New(&buf, "buf: ", log.LstdFlags)
	buflog.Println("hello")
	buflog.Print(" world")

	// Выводим содержимое буфера
	fmt.Print("from buflog:", buf.String())

	// Создаем JSON-хэндлер для slog и логгер, который использует этот хэндлер
	jsonhandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonhandler)
	myslog.Info("hi there")

	// Логирование с дополнительными ключ-значение парами
	myslog.Info("hello again", "key", "val", "age", 25)
}

//Вывод:
//standard logger
//2024/07/08 12:34:56.789012 with micro
//2024/07/08 12:34:56 main.go:21: with file/line
//2024/07/08 12:34:56 my: from mylog
//2024/07/08 12:34:56 ohmy: from mylog
//from buflog:2024/07/08 12:34:56 buf: hello
//{"time":"2024-07-08T12:34:56Z","level":"INFO","msg":"hi there"}
//{"time":"2024-07-08T12:34:56Z","level":"INFO","msg":"hello again","key":"val","age":25}

//Объяснение вывода
//Лог-сообщения от стандартного логгера, настроенные с различными флагами.
//Лог-сообщения от пользовательского логгера с различными префиксами.
//Лог-сообщения, записанные в буфер и затем выведенные на экран.
//Лог-сообщения в формате JSON, записанные с помощью slog.
