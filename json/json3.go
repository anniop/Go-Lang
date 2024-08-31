package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
	Age int64
	Location string
}

func main() {
	
	j:=[]byte(`{"name":"Ganesh","age":30,"location":"pune"}`)

	var p Person
	err := json.Unmarshal(j, &p)

	if err != nil {
		log.Fatalf("Unable to Decode")
	}
	fmt.Printf("Name:%s\nAge:%d\nLocation:%s\n", p.Name,p.Age,p.Location)
}
