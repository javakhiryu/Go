package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Определение подкоманды 'foo' и её флагов
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	fooEnable := fooCmd.Bool("enable", false, "can be true or false")
	fooName := fooCmd.String("name", "", "any string -name=example")

	// Определение подкоманды 'bar' и её флагов
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	barLevel := barCmd.Int("level", 0, "should be integer")

	// Проверка, что была передана хотя бы одна подкоманда
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
	}

	// Определение, какая подкоманда была вызвана
	switch os.Args[1] {
	case "foo":
		// Парсинг флагов подкоманды 'foo'
		fooCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'foo'")
		fmt.Println(" enable:", *fooEnable)
		fmt.Println(" name:", *fooName)
		fmt.Println(" tail:", fooCmd.Args())
	case "bar":
		// Парсинг флагов подкоманды 'bar'
		barCmd.Parse(os.Args[2:])
		fmt.Println("subcommand 'bar'")
		fmt.Println(" level:", *barLevel)
		fmt.Println(" tail:", barCmd.Args())

	default:
		// Сообщение об ошибке, если подкоманда не распознана
		fmt.Println("expected 'foo' or 'bar' subcommans")
		os.Exit(1)
	}
}
