package main

import "fmt"

func main() {

	myChannel := make(chan string)

	go func() {
	myChannel <- "Jay Ganesh"
	}()

	channel2 := <-myChannel
	fmt.Println(channel2)

}
