package main

import (
	"fmt"
)
func main(){
	var a = make(map[string] string)
	a["brand"] = "Ford"
	a["Model"] = "Mustang"
	a["Year"] = "1964"
	fmt.Printf(a["brand"])
}