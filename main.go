package main

import (
	"fmt"
	"go-with-tests/hello"
	"go-with-tests/integers"
)

func main() {
	// example 1: hello-world
	fmt.Println(hello.Hello("Shafee", "Spanish"))

	// example 2: integers
	fmt.Println(integers.Add(2, 5))
}
