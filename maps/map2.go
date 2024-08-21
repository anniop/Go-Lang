package main

import "fmt"

func main() {

	var MyMap = make(map[string]float64)

	MyMap["Ganesh"]=80.55
	MyMap["Krishna"]=90.76
	MyMap["Mahadev"]=74.52

	fmt.Println("MyMap : ",MyMap)
}
