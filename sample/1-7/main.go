package main

import (
	"fmt"

	"local.packages/hello"
)

func main() {
	name := hello.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}
