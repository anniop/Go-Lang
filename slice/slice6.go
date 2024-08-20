// Iterate over a slice using range in for loop

package main

import "fmt"

func main() {

	slice := []int {1,2,3,4,5,6,7,8,9,0}

	for index,value := range slice {
		fmt.Printf("Index = %d\t Value = %d\n",index,value)
	}
}
