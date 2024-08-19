/*Fifth Case :- In an array if the element type of the array is comparable then the array type is also comparable so we
can directly compare two arrays using == operator */

package main

import "fmt"

func main() {
	
	array1 := [5]int {1,2,3,4,5}
	array2 := [...] int {1,2,3,4,5}
	array3 := [5]int {6,3,7,2,8}

	fmt.Println(array1 == array2)
	fmt.Println(array2 == array3)
	fmt.Println(array1 == array3)
}
