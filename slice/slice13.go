// Appending one slice into another slice

package main

import "fmt"

func main() {

	slice1 := []int {1,2,3,4,5,6}
	slice2 := []int{10,20,30,40,50,60}

	slice2 = append(slice1,slice2...)
	fmt.Println(slice2)
}
