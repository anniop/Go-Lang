package main

import "fmt"

func main() {

	ch := make(chan string, 1)

	ch <- "Jay Ganesh"

	message := <-ch 
	fmt.Println("Message Received",message)

}
