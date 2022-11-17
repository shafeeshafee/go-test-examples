package main

import (
	"fmt"
	"go-with-tests/hello"
	"go-with-tests/integers"
	"go-with-tests/iteration"
)

// TESTING

func main() {
	// example 1: hello-world
	fmt.Println(hello.Hello("Shafee", "Spanish"))

	// example 2: integers
	fmt.Println(integers.Add(2, 5))

	// example 3: iteration
	fmt.Println(iteration.Repeat("u", 10))
}
