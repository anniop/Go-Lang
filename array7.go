package main

import "fmt"

func main() {

	intArray := [5] int {10,20,30,40,50}

	for index,value := range intArray {

		fmt.Println(index,"=>",value)
	}
}
