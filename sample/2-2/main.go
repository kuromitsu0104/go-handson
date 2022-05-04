package main

import (
	"fmt"
	"strconv"

	"github.com/kuromitsu0104/go-handson/hello"
)

func main() {
	x := hello.Input("type a price")
	n, err := strconv.Atoi(x)
	p := float64(n)
	fmt.Println(int(p * 1.1))

	// err変数が利用されていないのでエラー発生
	// 次のように呼び出せばOK
	fmt.Println(err)
}
