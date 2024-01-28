package main

import (
	"example/hello/monkey-interpreter/repl"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello! Welcome to the Monkey Programming Language!")
	repl.Start(os.Stdin, os.Stdout)
}
