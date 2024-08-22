package main

import "fmt"

type address struct {
	state string
	city string
	zipcode int
}

func main() {
	// Declaring and Initializing struct instance uninitialized fields are set to thir corresponding zero value
	a1 := address{state: "Maharashtra"}

	fmt.Println("Address1 : ",a1)

}
