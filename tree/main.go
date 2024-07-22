package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

func dirTree(out io.Writer, path string, printFiles bool) error {
	// Вложенная функция для печати дерева каталогов
	var printDirTree func(string, string, bool) error
	printDirTree = func(currentPath, indent string, last bool) error {
		entries, err := os.ReadDir(currentPath)
		if err != nil {
			return err
		}
		// Сортировка элементов по имени
		sort.Slice(entries, func(i, j int) bool {
			return entries[i].Name() < entries[j].Name()
		})

		for i, entry := range entries {
			// Пропустить файлы, если printFiles равно false
			if !printFiles && !entry.IsDir() {
				continue
			}

			// Определение префикса
			prefix := "├───"
			if last && i == len(entries)-1 {
				prefix = "└───"
			} else if !last && i == len(entries)-1 {
				prefix = "└───"
			}
			// Получение информации о файле
			info, err := entry.Info()
			if err != nil {
				return err
			}

			// Формирование строки с размером файла
			sizeStr := ""
			if !entry.IsDir() {
				if info.Size() == 0 {
					sizeStr = " (empty)"
				} else {
					sizeStr = fmt.Sprintf(" (%db)", info.Size())
				}
			}

			fmt.Fprintf(out, "%s%s%s%s\n", indent, prefix, entry.Name(), sizeStr)

			// Рекурсивный вызов для подкаталогов
			if entry.IsDir() {
				newIndent := indent + "│\t"
				if last && i == len(entries)-1 {
					newIndent = indent + "\t"
				} else if !last && i == len(entries)-1 {
					newIndent = indent + "\t"
				}
				err := printDirTree(filepath.Join(currentPath, entry.Name()), newIndent, i == len(entries)-1)
				if err != nil {
					return err
				}
			}
		}
		return nil
	}
	// Начать печать с указанного пути
	return printDirTree(path, "", false)
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
