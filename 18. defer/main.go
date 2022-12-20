package main

import "fmt"

func main() {
	// defer fmt.Println("World")
	// fmt.Println("Hello")
	fmt.Println("Start")
	panic("This is a panic")
	fmt.Println("End")
}