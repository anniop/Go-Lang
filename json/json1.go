package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {

	Name string `json:"First Name"`  
	Age int64
	Location string
}

func main() {
	
	person := Person{
		"Aniket", 21, "Mauje Sukene",
	}

	personArray,err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Unable to encode")

	}
	fmt.Println(string(personArray))

}
