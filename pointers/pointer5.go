package main

import "fmt"

func diplay(a *int) {
	
	*a = 500

}

func main() {

	var x = 100

	fmt.Printf("Value of x before function call : %d\n", x)
}
