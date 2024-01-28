package main

import (
	"example/hello/monkey-interpreter/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s! Welcome to the Monkey Programming Language!\n", user.Username)
	fmt.Println("Try typing in a Monkey command below :p")
	repl.Start(os.Stdin, os.Stdout)
}
