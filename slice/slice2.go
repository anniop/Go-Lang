// Create slice using an array.

package main

import "fmt"

func main() {
	array := [5]int {1,2,3,4,5}

	var slice1=array[1:2] // 2

	slice2:=array[0:] // 1,2,3,4,5

	slice3 := array[:2] // 1,2

	slice4 := array[:] // 1,2,3,4,5

	fmt.Println("Array : ",array)
	fmt.Println("Slice1 : ",slice1)
	fmt.Println("Slice2 : ",slice2)
	fmt.Println("Slice3 : ",slice3)
	fmt.Println("Slice4 : ",slice4)
}
