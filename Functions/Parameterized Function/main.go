package main

import "fmt"

func familyName(name string, age int) {
	fmt.Println("Hello I'm ", name, " And I am ", age, " years old")
}
func main() {
	familyName("Chirag", 20)
	familyName("Shiv", 20)
	familyName("Abhishek", 21)
}
