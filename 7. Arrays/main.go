package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}
	arr3 := [...]int{4, 5, 321, 124, 2}
	arr4 := [5]int{1:10, 2: 40};	
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr3[0])

	arr3[3] = 15
	fmt.Println(arr3[3])
}
