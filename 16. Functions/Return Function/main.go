package main
import "fmt"

func myFun(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + "World!"
	return
}

func main() {
	fmt.Println(myFun(5, "Hello"))
}
