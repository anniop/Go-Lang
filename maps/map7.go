package main

import (
	"fmt"
	"sort"
)

func main() {

	unsortedMap := map[int]string {
		50: "Ganesh",
		30: "Krishna",
		10: "Mahadev",
		40: "Vishnu",
		20: "Vasudev",
	}

	keys := make([]int,0,len(unsortedMap))

	for i:= range unsortedMap {
		keys = append(keys, i)
	}

	sort.Ints(keys)

	for _,value := range keys {
		fmt.Println(value,unsortedMap[value])
	}
}
