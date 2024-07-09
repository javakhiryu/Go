package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProg := os.Args

	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println("Args with program name: ", argsWithProg)
	fmt.Println("Args without progrma name: ", argsWithoutProg)
	fmt.Println("4th argument: ", arg)
}
