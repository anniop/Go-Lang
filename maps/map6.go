package main

import "fmt"

func main() {

	myMap := map[int]string {
		10: "Ganesh",
		20: "Krishna",
		30: "Mahadev",
		40: "Vishnu",
		50: "Vasudev",
	}

	for id,value:= range myMap {
		fmt.Printf("ID : %d \t Value : %s \n",id,value)
	}
}
