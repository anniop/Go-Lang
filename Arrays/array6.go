package main

import "fmt"

func main() {

	intArray := [5] int {1,2,3,4,5}

	for i:=0;i<len(intArray);i++ {
		fmt.Println(intArray[i])
	}
}
