package main

import "fmt"

func main() {

	intArray := [5] int {10,20,30,40,50}

	for _,value := range intArray {

		fmt.Println(value)
	}
}
