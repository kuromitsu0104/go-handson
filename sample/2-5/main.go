package main

import (
	"fmt"
	"strconv"

	"github.com/kuromitsu0104/go-handson/hello"
)

func main() {
	x := hello.Input("type a number")
	fmt.Print(x + "は、")
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数です。")
		} else {
			fmt.Println("奇数です。")
		}
	} else {
		fmt.Println("整数ではありません。")
	}
}
