package main

import "fmt"

type address struct {
	state string
	city string
	zipcode int
}

func main() {

    // Declaring and initializing struct instance using struct literal
	a1 := address{"Maharashtra", "Nashik", 422303}
	fmt.Println("Address1 : ",a1)

}
