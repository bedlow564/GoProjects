package main

import (
	"fmt"
	"go_cmd/hello"
	"os"
)

func main() {
	a := 2
	b := 2.1

	var c *int = &a

	fmt.Println(hello.Say(os.Args[1:]))

	fmt.Printf("a: %4T %[1]v\n", a) //[] bracket allows to shorten writing variable multipel times
	fmt.Printf("b: %4T %[1]v\n", b)
	fmt.Printf("Address of a is %v\n", c)
}
