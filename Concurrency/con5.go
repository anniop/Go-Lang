package main

import "fmt"


func main() {

	var channel1 chan int 
	fmt.Println("Channel1 Value:",channel1)
	fmt.Printf("Type of Channel :%T\n", channel1)

	channel2 := make(chan int)
	fmt.Println("Channel2 Value:",channel2)
	fmt.Printf("Type of Channel :%T\n", channel2)
}
