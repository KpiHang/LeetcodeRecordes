package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{3, 1, 2}
	slice2 := []string{"c", "a", "b"}

	sort.Slice(slice2, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})
	sort.Slice(slice1, func(i, j int) bool {
		return slice1[i] < slice1[j]
	})

	fmt.Println(slice1)
	fmt.Println(slice2)
}
