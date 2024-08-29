package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello to Main function") 
	
	go func() {
		fmt.Println("Jay Ganesh from go routine.............")
	}() 

	time.Sleep(1*time.Second)

	fmt.Println("Good Bye")
}
