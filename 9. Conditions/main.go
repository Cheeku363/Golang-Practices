package main

import "fmt"

func main() {
	var a int
	x := 30
	if x >= 10 {
		fmt.Println("x is larger than or equals to 10")
	} else if x > 20 {
		fmt.Println("x is larger than 20")
	} else {
		fmt.Println("x is less than 10")
	}
	fmt.Print(a)
}
