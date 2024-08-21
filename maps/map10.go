package main

import "fmt"

func main() {

	myMap := map[int]string {

		10: "Ganesh",
		20: "Krishna",
		30: "Mahadev",
	}

	for index, value := range myMap {
		fmt.Printf("Index : %d \t Value : %s \n", index, value)
	}

	myMap = make(map[int]string)

	fmt.Println(myMap)
}
