// Adding Elements to existing slice using append function

package main

import "fmt"

func main() {

	slice := make([]int,2,5)

	slice[0] = 10
	slice[1] = 20

	fmt.Println("Slice slice : ", slice)
	fmt.Printf("Length of slice : %d\t capacity of slice %d\n", len(slice), cap(slice))

	slice = append(slice, 30,40,50,60,70,80,90)

	fmt.Println("Slice after appending the data : ",slice)

	fmt.Printf("Length:%d\tCapacity:%d\t", len(slice),cap(slice))
}
