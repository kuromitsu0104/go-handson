package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name = input("type your name")
	fmt.Println("Hello, " + name + "!!")
}

func input(msg string) string {
	var scanner = bufio.NewScanner(os.Stdin)
	fmt.Print(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}
