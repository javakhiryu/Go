package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"b", "c", "a"}
	slices.Sort(strs)
	fmt.Println(strs)
	ints := []int{3, 2, 1}
	slices.Sort(ints)
	fmt.Println(ints)
	fmt.Println(slices.IsSorted(ints))
}
