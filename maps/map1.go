package main

import "fmt"

func main() {

	var emptymap = map[int]string { }
	fmt.Println("Empty Map :", emptymap)

	var myMap = map[int] string {
		1: "Ganesh",
		2: "Vishnu",
		3: "Mahadev",
		4: "Krishna",
	}

	fmt.Println("My Map :", myMap)
}
