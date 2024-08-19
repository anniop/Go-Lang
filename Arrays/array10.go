package main

import "fmt"

func main() {

	intArray := [5] int {10,20,30,40,50}

	j:=0
	for range intArray {
		fmt.Println(intArray[j])
		j++
	}
}
