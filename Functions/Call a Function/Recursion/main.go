package main

import "fmt"

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}

func factorial(x float64) (y float64) {
	if x > 0 {
		y = x * factorial(x-1)
	} else {
		y = 1
	}
	return
}
func main() {
	fact := factorial(5)
	fmt.Println(fact)
	testCount(1)
}
