package main

import (
	"flag"
	"fmt"
)

func main() {
	// Определение флагов командной строки
	// Флаг "word" типа string с значением по умолчанию "foo" и описанием "a string"
	wordPtr := flag.String("word", "foo", "a string")

	// Флаг "numb" типа int с значением по умолчанию 42 и описанием "an int"
	numPtr := flag.Int("numb", 42, "an int")

	// Флаг "fork" типа bool с значением по умолчанию false и описанием "a bool"
	forkPtr := flag.Bool("fork", false, "a bool")

	// Определение флага "svar" типа string с значением по умолчанию "bar" и описанием "a string var"
	// Здесь используется flag.StringVar для привязки значения к переменной svar
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Парсинг флагов командной строки
	flag.Parse()

	fmt.Println("word", *wordPtr) // Значение флага "word"
	fmt.Println("numb", *numPtr)  // Значение флага "numb"
	fmt.Println("fork", *forkPtr) // Значение флага "fork"
	fmt.Println("svar", svar)     // Значение флага "svar"

	// Вывод оставшихся аргументов после парсинга флагов
	fmt.Println("tail", flag.Args())

}

//go run commandLineFlags.go -word=hello -numb=123 -fork=true -svar=baz extra1 extra2
//word hello
//numb 123
//fork true
//svar baz
//tail [extra1 extra2]
