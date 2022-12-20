package main

import "fmt"

type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person
	// Pers1 specification
	pers1.name = "Chirag"
	pers1.age = 20
	pers1.job = "Software Developer"
	pers1.salary = 5000

	// Pers2 specification
	pers2.name = "Shiv"
	pers2.age = 22
	pers2.job = "Student"
	pers2.salary = 6500

	// Access and Print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and Print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)
}
