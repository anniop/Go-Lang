// Iterate over a slice using for loop as while loop


package main
import "fmt"

func main() {

	slice := []int {1,2,3,4,5,6,7,8,9,0}

	i:=0
	for i < len(slice) {
		fmt.Printf("value = %d\n", slice[i])
		i++
	}
}
