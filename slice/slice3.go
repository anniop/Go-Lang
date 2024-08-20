// Create Slice using already existing Slice

package main

import "fmt"

func main() {
	 mainslice := [] int {1,2,3,4,5,6,7}
	
	var slice1 =mainslice[1:5] // 2,3,4,5  
	slice2 := mainslice[2:] // 3,4,5,6,7
	slice3 := mainslice[:4] // 1,2,3,4

	fmt.Println("Main Slice : ",mainslice)
	fmt.Println("Silce1 : ",slice1)
	fmt.Println("Slice2 : ", slice2)
	fmt.Println("Slice3 : ",slice3)
	
}
