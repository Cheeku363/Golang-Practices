package main
import "fmt"

func main() {
	// With var keyword
	var var1 int = 100 //type is integer
	var var3 = 100     // type is automatically decided

	// With := sign
	var2 := 10 // type is automatically decided

	// MultiLine Variable
	var a, b, c, d int = 1, 3, 5, 7
	var e, f = 6, "Hello"
	g, h := 10, "World!"

	// Go variable declaration in a block
	var (
		x int
		y int = 1
		z string =  "Hello"
	) 
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(f)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
