package main

import "fmt"

func main() {
	var a = make(map[string]string) // the map is empty now
	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"

	fmt.Printf("a\t%v\n", a)
}
