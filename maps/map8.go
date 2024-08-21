package main

import (
	"fmt"
	"sort"
)

func main() {

	unsortedMap := map[int]string{
		50: "Ganesh",
		30: "Krishna",
		10: "Mahadev",
		40: "Vishnu",
		20: "Vasudev",
	}
	values := make([]string,0,len(unsortedMap))

	for _,value := range unsortedMap {
		values = append(values,value)
	}
	sort.Strings(values)

	for _,value := range values {
		fmt.Println(value)
	}
}
