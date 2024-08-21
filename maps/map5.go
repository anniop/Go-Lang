package main

import "fmt"

func main() {

	myMap := map[int]string{
		10: "Ganesh",
		20: "Krishna",
		30: "Mahadev",
		40: "Vishnu",
		50: "Vasudev",
	}
	fmt.Println("MyMap : ",myMap)

	myMap[60] = "Bhrama"
	myMap[70] = "Ganpati"

	fmt.Println("Map After adding elements : ",myMap)


	myMap[20] = "Mahadev"
	myMap[30] = "Krishna"

	fmt.Println("map after updating : ",myMap)

	delete(myMap,70)
	fmt.Println("map after deleting : ",myMap)

}
