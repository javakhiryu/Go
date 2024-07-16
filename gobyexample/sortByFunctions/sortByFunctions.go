//Ваш код демонстрирует использование различных методов сортировки в Go, включая sort.Slice для сортировки срезов и slices.SortFunc для сортировки срезов с использованием функции сравнения из пакета slices. Давайте рассмотрим ваш код с подробными комментариями и объяснением каждого элемента:

package main

import (
	"cmp" // Пакет для сравнения элементов (доступен в Go 1.21+)
	"fmt"
	"slices" // Пакет для работы с срезами (доступен в Go 1.18+)
	"sort"   // Пакет для сортировки срезов и других коллекций
)

func main() {
	// Срез строк для сортировки
	fruits := []string{"banana", "apple", "lemon"}

	// Используем sort.Slice для сортировки среза строк в алфавитном порядке
	sort.Slice(fruits, func(i, j int) bool {
		return fruits[i] < fruits[j]
	})

	// Выводим отсортированный срез строк
	fmt.Println(fruits) // ["apple", "banana", "lemon"]

	// Определяем тип Person с полями name и age
	type Person struct {
		name string
		age  int
	}

	// Создаем срез с несколькими Person
	persons := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}

	// Используем slices.SortFunc для сортировки среза persons по возрасту
	slices.SortFunc(persons, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	// Выводим отсортированный срез Person
	fmt.Println(persons) // [{TJ 25} {Jax 37} {Alex 72}]
}

//В Go 1.18 и выше доступны новые функции и пакеты для работы со срезами и коллекциями:
//
//slices.Sort(slice): Сортирует срез в порядке возрастания. В вашем коде вы не использовали эту функцию, но это полезная функция для сортировки срезов.
//slices.IsSorted(slice): Проверяет, отсортирован ли срез.
//slices.SortFunc(slice, less func(a, b T) bool): В вашем коде эта функция сортирует persons по возрасту.
