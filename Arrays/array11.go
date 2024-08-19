package main

import "fmt"

func main() {

	array1 := [3]string {"aniket","ganesh","jay"}

	array2 := array1

	array3 := &array1

	fmt.Printf("array1 : %v\n",array1)
	fmt.Printf("array2 : %v\n", array2)

	array1[0]= "Varun"

	fmt.Printf("array1 : %v\n",array1)
	fmt.Printf("array2 : %v\n", array2)
	
	fmt.Printf("Array3 : %v\n", *array3)
}
