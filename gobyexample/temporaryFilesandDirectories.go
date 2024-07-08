package main

import (
	"fmt"           // Пакет для форматированного ввода-вывода
	"os"            // Пакет для работы с операционной системой
	"path/filepath" // Пакет для работы с путями файловой системы
)

// Функция для проверки ошибок
func check(err error) {
	if err != nil {
		panic(err) // Если ошибка не равна nil, вызвать панику
	}
}

func main() {

	// Создание временного файла с префиксом "sample"
	f, err := os.CreateTemp("", "sample")
	check(err)

	// Вывод имени временного файла
	fmt.Println("Temp file name: ", f.Name())

	// Отложенное закрытие файла и его удаление после завершения main
	defer func() {
		f.Close() //Когда файл открыт, операционная система может блокировать его удаление. Вот обновленный код, в котором добавлено закрытие файла перед удалением:
		os.Remove(f.Name())
	}()

	// Запись байтов {1, 2, 3, 4} в временный файл
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	// Создание временной директории с префиксом "sampledir"
	dname, err := os.MkdirTemp("", "sampledir")
	check(err)

	// Вывод имени временной директории
	fmt.Println("Temp dir name: ", dname)

	// Отложенное удаление временной директории после завершения main
	defer os.RemoveAll(dname)

	// Создание пути для файла "file1" внутри временной директории
	fname := filepath.Join(dname, "file1")

	// Запись байтов {1, 2} в файл "file1" с правами доступа 0666
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
