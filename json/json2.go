package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {

	Name string `json:"First Name"`  
	Age int64
	Location string  `json:"Location,omitempty"`
}

func main() {
	

	person := Person{
		Name: "Ganesh",
		Age : 30,
	}
	personArray,err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Unable to encode")

	}
	fmt.Println(string(personArray))

}
