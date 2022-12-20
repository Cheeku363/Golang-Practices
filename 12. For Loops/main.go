package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for i := 0; i < 7; i++ {
		if i == 3 {
			continue
		}
		if i == 7{
			break
		}
		fmt.Println(i)
	}

}
