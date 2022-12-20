package main

import (
	"fmt"
	"time"
)


func main(){
	message := "Hello user"
	go func(){
		fmt.Println(message)
	}()
	message = "Hello youtube"
	time.Sleep(100 * time.Millisecond)
}
