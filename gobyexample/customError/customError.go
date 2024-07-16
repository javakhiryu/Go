//Ваш код демонстрирует, как создавать и использовать пользовательские ошибки в Go.
//Вы определяете структуру argError для представления ошибки, которая включает аргумент и сообщение, а затем используете её в функции f.
//В main вы обрабатываете ошибку, проверяя, является ли она типом argError.

package main

import (
	"errors"
	"fmt"
)

// Определение пользовательской ошибки
type argError struct {
	arg     int
	message string
}

// Реализация метода Error для интерфейса error
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

// Функция f возвращает ошибку argError, если аргумент равен 42
func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	// Вызов функции f с аргументом 42
	_, err := f(42)
	var ae *argError

	// Проверка, является ли ошибка типом argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)     // Вывод аргумента, вызвавшего ошибку
		fmt.Println(ae.message) // Вывод сообщения об ошибке
	} else {
		fmt.Println("error does not match argError")
	}
}

//Объяснение:
//
//Определение argError: Структура argError содержит два поля: arg и message.
//Метод Error: Метод Error реализует интерфейс error и возвращает строковое представление ошибки.
//Функция f: Эта функция принимает целое число arg и возвращает либо результат вычислений, либо ошибку argError, если arg равен 42.
//Проверка ошибки в main:
//	Вызывается функция f с аргументом 42.
//	Используется errors.As для проверки, является ли ошибка типом argError.
//	Если это так, выводятся аргумент и сообщение ошибки. В противном случае выводится сообщение о несоответствии типа ошибки.
