// Iterate over a slice using a blank identifier in for loop

package main

import "fmt"

func main() {

	slice := []int {1,2,3,4,5,6,7,8,9,0}

	for _,value := range slice {
		fmt.Printf("Value = %d\n",value)
	}
}
