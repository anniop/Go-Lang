// Fourth Case :- In an array you are allowed to iterate over the range of the elements of the array

package main

import "fmt"

func main() {

	array := [...]int {10,20,30,40,50,60,70,80,90}

	for i:=0;i<len(array);i++ {
		fmt.Println("Value of array is :", array[i])
	}
}
