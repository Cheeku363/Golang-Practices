package main

import "fmt"

func main() {
	// Print Statement
	var i, j string = "Hello", "World!"
	// fmt.Print(i)
	// fmt.Print(j)
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	// Printf statement
	var x string = "Hello"
	var y int = 15
	fmt.Printf("x has the value: %v and type: %T\n", x, x)
	fmt.Printf("y has the value: %v and type: %T\n", y, y)
}
