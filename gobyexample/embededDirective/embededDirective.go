package main

import (
	"embed"
	"fmt"
	"log"
)

// Встраивание файла "folder/single_file.txt" как строку
//
//go:embed folder/single_file.txt
var fileString string

// Встраивание файла "folder/single_file.txt" как байтовый срез
//
//go:embed folder/single_file.txt
var fileByte []byte

// Встраивание всех файлов с расширением .hash из папки "folder" и файл "folder/single_file.txt" в файловую систему
//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
var folder embed.FS

func main() {

	// Печать содержимого "folder/single_file.txt" как строки
	print(fileString)

	// Печать содержимого "folder/single_file.txt" как байтового среза, преобразованного в строку
	print(string(fileByte))

	// Чтение и печать содержимого "folder/file1.hash" из встраиваемой файловой системы
	// Чтение и печать содержимого "folder/file2.hash" из встраиваемой файловой системы
	if content1, err := folder.ReadFile("folder/file1.hash"); err != nil {
		log.Printf("Error reading file1.hash: %v\n", err)
	} else {
		fmt.Println("Content of file1.hash:")
		fmt.Println(string(content1))
	}

	// Проверка наличия файла "folder/file2.hash" и его вывод
	if content2, err := folder.ReadFile("folder/file2.hash"); err != nil {
		log.Printf("Error reading file2.hash: %v\n", err)
	} else {
		fmt.Println("Content of file2.hash:")
		fmt.Println(string(content2))
	}

	// Список всех файлов в встроенной файловой системе
	if entries, err := folder.ReadDir("folder"); err != nil {
		log.Printf("Error reading directory 'folder': %v\n", err)
	} else {
		fmt.Println("Files in embedded folder:")
		for _, entry := range entries {
			fmt.Println(" -", entry.Name())
		}
	}
}
