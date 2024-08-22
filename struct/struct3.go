package main

import "fmt"

type address struct {
	state string
	city string
	zipcode int
}

func main() {
   // Declaring and initializing struct instance using naming fields
   a1 := address{state: "Maharashtra", city : "Nashik", zipcode: 422302}
   fmt.Println("Address1 :", a1)
	

}
