package main

import "fmt"

type address struct {
	state string
	city string
	zipcode int
}

func main() {

    // Struct instantiation using new keyword
	a1 := new(address)
	a1.state = "Maharashtra"
	a1.city = "Nashik"
	a1.zipcode = 422302

	fmt.Println("Address1 : ",a1)

}
