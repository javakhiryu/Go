package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args // Получаем все аргументы, включая имя программы


	argsWithoutProg := os.Args[1:] // Получаем все аргументы, кроме имени программы

	// Получаем четвертый аргумент (индексация начинается с 0)
	// Обратите внимание, что программа завершится ошибкой, если четвертый аргумент отсутствует.
	arg := os.Args[3]

	fmt.Println("Args with program name: ", argsWithProg) // Выводим все аргументы, включая имя программы
	fmt.Println("Args without progrma name: ", argsWithoutProg) // Выводим все аргументы, кроме имени программы
	fmt.Println("4th argument: ", arg) // Выводим четвертый аргумент

//go run main.go arg1 arg2 arg3 arg4 arg5

//Args with program name:  [main.go arg1 arg2 arg3 arg4 arg5]
//Args without program name:  [arg1 arg2 arg3 arg4 arg5]
//4th argument:  arg4

