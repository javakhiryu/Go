package main

import (
	"fmt"           // Пакет для форматированного ввода-вывода
	"path/filepath" // Пакет для работы с файловыми путями
	"strings"       // Пакет для работы со строками
)

func main() {

	// Объединяем несколько частей пути в один
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p: ", p) // Результат: "dir1/dir2/filename" на Unix или "dir1\dir2\filename" на Windows

	// Пример с лишними слэшами
	fmt.Println(filepath.Join("dir1//", "filename"))       // Результат: "dir1/filename"
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // Результат: "dir1/filename" после нормализации пути

	// Получение директории из пути
	fmt.Println("Dir(p):", filepath.Dir(p)) // Результат: "dir1/dir2" на Unix или "dir1\dir2" на Windows

	// Получение базового имени файла из пути
	fmt.Println("Base(p):", filepath.Base(p)) // Результат: "filename"

	// Проверка, является ли путь абсолютным
	fmt.Println(filepath.IsAbs("dir/file"))  // Результат: false (относительный путь)
	fmt.Println(filepath.IsAbs("/dir/file")) // Результат: true на Unix, false на Windows

	// Работа с расширением файла
	filename := "config.json"

	// Получение расширения файла
	ext := filepath.Ext(filename)
	fmt.Println(ext) // Результат: ".json"

	// Удаление расширения из имени файла
	fmt.Println(strings.TrimSuffix(filename, ext)) // Результат: "config"

	// Вычисление относительного пути
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err) // Если ошибка, паника
	}
	fmt.Println(rel) // Результат: "t/file"

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err) // Если ошибка, паника
	}
	fmt.Println(rel) // Результат: "../c/t/file"
}

//Этот код демонстрирует основные функции пакета path/filepath, такие как объединение частей пути, получение директории и базового имени файла, проверка абсолютности пути, работа с расширениями файлов и вычисление относительного пути.
