package main

import (
	"fmt"
	"time"
)

func name () {

	array1 := [4]string{"Jay","Ganesh","Mahadev","vishnu"}

	for i:= 0; i < len(array1); i++ {
		time.Sleep(150*time.Millisecond)
		fmt.Printf("%s\n", array1[i])
	}
}

func id () {
	array2 := [4]int {100,101,102,103}
	for i:=0;i<len(array2);i++{
		time.Sleep(500*time.Millisecond)
		fmt.Printf("%d\n", array2[i])
	}
}

func main() {

	fmt.Println("Main Goroutine Start................")
	
	go name()

	go id()
	time.Sleep(1500*time.Millisecond)
	fmt.Println("Main Goroutine End...................")
}

