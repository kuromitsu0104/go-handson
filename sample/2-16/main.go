package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := [5]int{10, 20, 30, 40, 50}
	b := a[0:3]
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	fmt.Println(b)
	fmt.Println(reflect.TypeOf(b))
}
