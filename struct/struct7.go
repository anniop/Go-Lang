package main

import "fmt"

type address struct {
	state string
	city string
	zipcode int
}

func main() {

	// Struct instantiation using pointer address operator
	var a = &address{}
	a.city = "Nashik"
	a.state = "Maharashtra"
	a.zipcode = 422302
	fmt.Println("Address1 : ",a)

}
