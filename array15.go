// Second Case :- In an array you can find the length of the array using len() function.

package main

import "fmt"

func main() {

	array1 := [3] int {1,2,3}
	array2 := [...] int {1,2,3,4,5}

	fmt.Println("Length of the array1 is : ",len(array1))
	fmt.Println("Length of the array2 is : ",len(array2))
}
