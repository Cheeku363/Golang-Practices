package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"] = "Ford"
	a["Model"] = "Mustang"
	a["Year"] = "1964"
	fmt.Println(a)

	a["Year"] = "1979" // update an element
	a["color"] = "red" // added an element
	fmt.Println(a)
}
