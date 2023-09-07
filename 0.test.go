package main

import "fmt"

func main() {
	nums := []int{3, 34, 5, 5, 6, 7, 7}
	for i, x := range nums[1:5] {
		fmt.Println(i, x)
	}
}
