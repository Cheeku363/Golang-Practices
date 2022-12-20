package main

import "fmt"

type Foo struct{
	bar int
}

func main(){
	var foo *Foo
	foo = new (Foo)
	fmt.Println(foo)
	
	// var a int = 12
	// var b *int = &a
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(*b)
}
