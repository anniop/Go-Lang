package main

import "fmt"

func main() {

	var x int = 0

	for x < 12 {

		if x ==5 {
			x = x + 2
			continue
		}

		fmt.Printf("value of x is %d\n", x)
		x++
	}
}
