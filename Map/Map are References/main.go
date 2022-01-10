package main

import "fmt"

func main() {
	var a = map[string]string{"brand": "ford", "model": "Mustang", "year": "1964"}
	b := a //Referencing
	fmt.Println(a)
	fmt.Println(b)
	b["year"] = "1970"
	fmt.Println("After change to b: ")
	fmt.Println(a) // both store same changes
	fmt.Println(b) // both store same changes
}
