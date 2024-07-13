package main

import(
	"fmt"
	"os"
)

func main(){

	// Эта функция не будет вызвана из-за os.Exit(3)
	defer fmt.Println("!")

	// Завершает выполнение программы с кодом возврата 3
	os.Exit(3)
}