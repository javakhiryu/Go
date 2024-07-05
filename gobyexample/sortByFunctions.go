package main

import (
	"cmp"
	"fmt"
	"slices"
	"sort"
)

func main() {
	fruits := []string{"banana", "apple", "lemon"}

	sort.Slice(fruits, func(i, j int) bool {
		return fruits[i] < fruits[j]
	})

	fmt.Println(fruits)

	type Person struct {
		name string
		age  int
	}
	persons := []Person{
		Person{name: "Jax", age: 37},
		Person{name: "TJ", age: 25},
		Person{name: "Alex", age: 72},
	}
	slices.SortFunc(persons, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})
	fmt.Println(persons)
}
