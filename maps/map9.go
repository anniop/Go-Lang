package main

import "fmt"


func main() {

	myMap := map[int]string {

		10: "Ganesh",
		20: "Krishna",
		30: "Mahadev",
	}

	for index ,value := range myMap{
		fmt.Printf("Index : %d\t Value : %s \n",index,value)
	}

	for i := range myMap {
	delete(myMap, i)
	}

	fmt.Println("My Map After Deleting : ")
	fmt.Println(myMap)
}
