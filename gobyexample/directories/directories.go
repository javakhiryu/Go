package main

import (
	"fmt"           // Пакет для форматированного ввода-вывода
	"io/fs"         // Пакет для интерфейсов файловой системы
	"os"            // Пакет для работы с операционной системой
	"path/filepath" // Пакет для работы с путями файловой системы
)

// Функция для проверки ошибок
func check(e error) {
	if e != nil {
		panic(e) // Если ошибка не равна nil, вызвать панику
	}
}

func main() {
	// Создание директории "subdir" с правами доступа 0755
	err := os.Mkdir("subdir", 0755)

	check(err)

	// Отложенное удаление всей директории "subdir" после завершения main
	defer os.RemoveAll("subdir")

	// Анонимная функция для создания пустого файла
	createEmptyFile := func(name string) {
		d := []byte("")                    // Создание пустого байтового среза
		check(os.WriteFile(name, d, 0644)) // Создание файла с правами доступа 0644
	}

	// Создание пустого файла "subdir/file1"
	createEmptyFile("subdir/file1")

	// Создание вложенных директорий "subdir/parent/child" с правами доступа 0755
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	// Создание нескольких пустых файлов в разных поддиректориях
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// Чтение содержимого директории "subdir/parent"
	c, err := os.ReadDir("subdir/parent")
	check(err)

	fmt.Println("Listing subdir/parent")

	// Вывод списка содержимого директории "subdir/parent"
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Смена текущей директории на "subdir/parent/child"
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Чтение содержимого текущей директории
	c, err = os.ReadDir(".")
	check(err)

	fmt.Println("Listing subdir/parent/child")

	// Вывод списка содержимого текущей директории
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Возврат в корневую директорию проекта
	err = os.Chdir("../../..")
	check(err)

	fmt.Println("Visiting subdir")
	// Рекурсивный обход директории "subdir" и вывод её содержимого
	err = filepath.WalkDir("subdir", visit)
	check(err)
}

// Функция для обработки каждого файла или директории при обходе
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err // Если ошибка, вернуть её
	}

	fmt.Println(" ", path, d.IsDir()) // Вывод пути и типа (директория или файл)
	return nil
}

//Вывод:

//Listing subdir/parent
//  child true
//  file2 false
//  file3 false
//Listing subdir/parent/child
//  file4 false
//Visiting subdir
//  subdir true
//  subdir\file1 false
//  subdir\parent true
//  subdir\parent\child true
//  subdir\parent\child\file4 false
//  subdir\parent\file2 false
//  subdir\parent\file3 false
