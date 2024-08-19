package main

import "fmt"

func main() {

	var x int = 0

	label1:

	for x < 8 {
		if (x == 5){
			x = x + 1
			goto label1
		}
		fmt.Println("Value of x : %d\n",x)
		x++
	}
}
