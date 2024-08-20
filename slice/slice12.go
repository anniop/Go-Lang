// Copy Slice items to another Slice

package main

import "fmt"

func main () {

	a := []int {10,20,30,40,50,60,70,80,90,100}

	fmt.Println("Slice A : ",a)
	fmt.Printf("Length : %d\t Capacity : %d\n",len(a),cap(a))

	b := make([]int,10,10)
	copy(b,a)

	fmt.Println("Slice B : ",b)
	fmt.Printf("Length : %d\t Capacity : %d\n",len(b),cap(b))
}
