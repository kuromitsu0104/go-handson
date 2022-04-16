package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(pop(a))
}

func push(a []int, v int) []int {
	return append(a, v)
}

func pop(a []int) []int {
	return a[:len(a)-1]
}
