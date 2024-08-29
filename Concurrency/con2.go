package main

import (
	"fmt"
	"time"
)

func main() {

	go hello1()
	fmt.Println("Jay Ganesh")

	time.Sleep(1*time.Second)
	fmt.Println("Good Bye")

}

func hello1() {
	go hello2()
	fmt.Println("Jay Ganesh from Goroutine 1")
}

func hello2() {
	fmt.Println("Jay Ganesh from  Goroutine 2")
}


