package main
import "fmt"
func main(){
	day := 8
	switch day{
	case 1: fmt.Println("Mon")
	case 2: fmt.Println("Tue") // In Golang we doesn't need to specify break keyword.
	case 3: fmt.Println("Wed")
	case 4: fmt.Println("Thurs")
	case 5: fmt.Println("Fri")
	case 6: fmt.Println("Sat")
	case 7: fmt.Println("Sun")
	default: fmt.Println("not a week day")
	}

	days := 5
	switch days {
	case 1, 3, 5: fmt.Println("Odd weekdays")
	case 2, 4: fmt.Println("Even weekdays")
	case 6, 7: fmt.Println("Weekend")
	default: fmt.Println("Invalid day of day number")
	}
}