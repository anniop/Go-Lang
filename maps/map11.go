package main

import "fmt"

func main() {

	first := map[int]string {

		10: "Ganesh",
		20: "Krishna",
		30: "Mahadev",
	}

	second := map[int]string {

		40: "Vishnu",
		50: "Vasudev",
		60: "bhrama",
	}

	for i,j := range second {
		first[i] = j
	}

	for i,j := range first {
		fmt.Printf("Index : %d\t Value : %s\n", i,j)
	}
}
