package main

import (
	"fmt"
)

func main() {

	ch := make(chan string)
	
	go send(ch)

	for {
		result,ok := <-ch

		if !ok {
			fmt.Println("Channel Closed", ok)
			 break
		}
		fmt.Println("Channel Open", result,ok)
	}
}

func send(ch chan string) {

	for v:=0;v<4;v++{
		ch <- "Jay Ganesh"
	}

	close(ch)
}
