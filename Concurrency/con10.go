package main

import "fmt"

func main() {

	myChannel := make(chan string)
	myChannel2 := make(chan string)
	go func() {
		myChannel <- "Jay Ganesh"
	}()
	
	go func() {
		myChannel2 <- "Har Har Mahadev"
	}()

	select {
	case channel1 := <-myChannel:
			fmt.Println(channel1)
	case channel2 := <-myChannel2:
			fmt.Println(channel2)
	}

}
