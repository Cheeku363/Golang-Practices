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

	printPerson(pers1)
	printPerson(pers2)
}

func printPerson(pers Person){
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}