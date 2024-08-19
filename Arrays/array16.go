/* Third Case :- In an array if ellipsis "..." become visible at the place of length then the length of the array is
determined by the initialized elements. */


package main

import "fmt"

func main() {

	array := [...]string {"C","C++","GoLang","Java"}

	fmt.Println("The Elements in the arry are : ",array)
	fmt.Println("The Length of the string is : ", len(array))
}
