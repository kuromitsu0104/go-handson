package main

import (
	"fmt"

	"github.com/kuromitsu0104/go-handson/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
