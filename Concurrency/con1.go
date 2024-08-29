package main

import (
	"fmt"
	"time"
)
func main() {

	go display()
	fmt.Println("Hello")
	time.Sleep(1*time.Second)
	fmt.Println("GoodBye")

}

func display() {
	fmt.Println("Jay Ganesh From GoRoutine")
}
