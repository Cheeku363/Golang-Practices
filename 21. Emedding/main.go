package main
import "fmt"

type Processor struct{
	processorName string
	cores int
}
type Memory struct{
	memoryCapacity int
	memoryType string
}

// has a relationship
type Computer struct{
	Processor
	Memory
	price int
}

func main(){
	computer := Computer{}
	computer.price = 5000;
	computer.processorName = "Intel i5"
	computer.cores = 6
	computer.memoryCapacity = 8
	computer.memoryType = "DDR4"
	fmt.Println(computer)
}